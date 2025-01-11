package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/api/handlers"
)

func adminRoutes(r *gin.Engine, adminHandler *handlers.AdminHandler) {
	r.POST("/v1/admin/create", adminHandler.CreateAdmin)
}
