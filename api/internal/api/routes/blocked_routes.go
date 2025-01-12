package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/api/handlers"
	"github.com/valu/vemeet-admin-api/internal/api/middleware"
)

func blockedRoutes(r *gin.Engine, blockedHandlers *handlers.BlockedHandler) {
	b := r.Group("/v1/blocked")
	b.Use(middleware.RequireAuthenticatedUser())
	{
		b.GET("", blockedHandlers.GetBlockeds)
		b.GET("/:id", blockedHandlers.GetBlockedById)
		b.PATCH("/:id", blockedHandlers.UpdateBlocked)
		b.POST("", blockedHandlers.CreateBlocked)
	}
}
