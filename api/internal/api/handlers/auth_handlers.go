package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/errors"
	"github.com/valu/vemeet-admin-api/internal/models"
	"github.com/valu/vemeet-admin-api/internal/services"
)

type AuthHandler struct {
	authService services.AuthServiceInterface
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input models.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.HandleError(c, err)
		return
	}

	tokens, admin, err := h.authService.LoginUser(input.Email, input.Password)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
		"admin":  admin,
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("Refresh-Token-X")
	if refreshToken == "" {
		errors.HandleError(c, errors.NewAuthenticationError("refresh token is required"))
		return
	}

	tokens, err := h.authService.RefreshTokens(refreshToken)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}

func (h *AuthHandler) Session(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		errors.HandleError(c, errors.NewAuthenticationError("user id not found"))
		return
	}

	userId, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		errors.HandleError(c, errors.NewAuthenticationError("invalid user id"))
		return
	}

	admin, err := h.authService.GetSession(userId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admin": admin,
	})
}
