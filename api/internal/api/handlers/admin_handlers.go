package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/data"
	"github.com/valu/vemeet-admin-api/internal/errors"
	"github.com/valu/vemeet-admin-api/internal/models"
	"github.com/valu/vemeet-admin-api/internal/services"
)

type AdminHandler struct {
	adminService services.AdminServiceInterface
}

func NewAdminHandler(adminService services.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{adminService}
}

func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var input models.CreateAdminRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		errors.HandleError(c, err)
		return
	}

	admin := data.Admin{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	}
	fmt.Println(admin)
	err := h.adminService.InsertAdmin(&admin)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, admin)
}
