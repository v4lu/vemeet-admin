package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/errors"
	"github.com/valu/vemeet-admin-api/internal/services"
)

type UserHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.NewValidationError("invalid user id"))
		return
	}

	user, err := h.userService.GetUserById(id)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
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

	users, err := h.userService.GetUsers(page, limit, sort, order, search)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
