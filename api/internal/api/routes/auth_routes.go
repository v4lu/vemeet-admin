package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/api/handlers"
	"github.com/valu/vemeet-admin-api/internal/api/middleware"
)

func authRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	a := r.Group("/v1/auth")

	a.POST("/login", authHandler.Login)
	a.POST("/refresh-token", authHandler.RefreshToken)
	a.GET("/", middleware.RequireAuthenticatedUser(), authHandler.Session)
}
