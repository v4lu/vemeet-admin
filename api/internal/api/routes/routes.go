package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/valu/vemeet-admin-api/internal/api/handlers"
	"github.com/valu/vemeet-admin-api/internal/api/middleware"
	"github.com/valu/vemeet-admin-api/internal/auth"
)

type Router struct {
	adminHandler *handlers.AdminHandler
	authHandler  *handlers.AuthHandler
	userHandler  *handlers.UserHandler
	tokenManager *auth.TokenManager
	router       *gin.Engine
}

func NewRouter(
	r *gin.Engine,
	adminHandler *handlers.AdminHandler,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	tokenManager *auth.TokenManager,
) *Router {
	return &Router{adminHandler, authHandler, userHandler, tokenManager, r}
}

func (r *Router) Router() {
	r.router.Use(gin.Recovery())
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.router.Use(middleware.AuthMiddleware(r.tokenManager))
	adminRoutes(r.router, r.adminHandler)
	authRoutes(r.router, r.authHandler)
	userRoutes(r.router, r.userHandler)
}
