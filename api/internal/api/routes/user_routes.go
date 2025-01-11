package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/valu/vemeet-admin-api/internal/api/handlers"
	"github.com/valu/vemeet-admin-api/internal/api/middleware"
)

func userRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	u := r.Group("/v1/users")
	u.Use(middleware.RequireAuthenticatedUser())
	{
		u.GET("", userHandler.GetUsers)
		u.GET("/:id", userHandler.GetUserById)
		u.GET("/username/:username", userHandler.GetUserByUsername)
	}
}
