package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/data"
	"github.com/valu/vemeet-admin-api/internal/errors"
	"github.com/valu/vemeet-admin-api/internal/services"
)

type BlockedHandler struct {
	blockedService services.BlockedServiceInterface
	userService    services.UserServiceInterface
}

func NewBlockedHandler(
	blockedService services.BlockedServiceInterface,
	userService services.UserServiceInterface,
) *BlockedHandler {
	return &BlockedHandler{
		blockedService: blockedService,
		userService:    userService,
	}
}

func (h *BlockedHandler) GetBlockedById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid blocked id"))
		return
	}

	blocked, err := h.blockedService.GetBlockedById(id)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, blocked)
}

func (h *BlockedHandler) GetBlockeds(c *gin.Context) {
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid page"))
		return
	}

	limit, err := strconv.ParseInt(c.DefaultQuery("pageSize", "10"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid limit"))
		return
	}

	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")
	search := c.DefaultQuery("search", "")

	blockeds, err := h.blockedService.GetBlockeds(page, limit, sort, order, search)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, blockeds)
}

type BlockedCreateRequest struct {
	UserID int64  `json:"user_id"`
	Reason string `json:"reason"`
}

func (h *BlockedHandler) CreateBlocked(c *gin.Context) {
	var blocked BlockedCreateRequest
	if err := c.ShouldBindJSON(&blocked); err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid blocked data"))
		return
	}

	blockedReq := &data.Blocked{
		UserID: blocked.UserID,
		Reason: blocked.Reason,
	}

	createdBlocked, err := h.blockedService.CreateBlocked(blockedReq)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	_, err = h.userService.ToggleBlockUser(blocked.UserID)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createdBlocked)
}

type BlockedUpdateRequest struct {
	Reason string `json:"reason"`
}

func (h *BlockedHandler) UpdateBlocked(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid blocked id"))
		return
	}

	var blocked BlockedUpdateRequest
	if err := c.ShouldBindJSON(&blocked); err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid blocked data"))
		return
	}

	blockedReq := &data.Blocked{
		ID:     id,
		Reason: blocked.Reason,
	}

	updatedBlocked, err := h.blockedService.UpdateBlocked(blockedReq)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedBlocked)
}

func (h *BlockedHandler) DeleteBlocked(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid blocked id"))
		return
	}

	deleted, err := h.blockedService.DeleteBlocked(id)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	user, err := h.userService.GetUserById(id)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if user.Blocked {
		_, err := h.userService.ToggleBlockUser(id)
		if err != nil {
			errors.HandleError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"deleted": deleted})
}
