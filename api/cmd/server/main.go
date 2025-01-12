// @title           Vemeet ADMIN API
// @version         1.0
// @description     A REST API service for managing Vemeet ADMIN dashboard
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9001
// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/valu/vemeet-admin-api/docs"
	"github.com/valu/vemeet-admin-api/internal/api/handlers"
	"github.com/valu/vemeet-admin-api/internal/api/routes"
	"github.com/valu/vemeet-admin-api/internal/auth"
	"github.com/valu/vemeet-admin-api/internal/config"
	"github.com/valu/vemeet-admin-api/internal/data"
	"github.com/valu/vemeet-admin-api/internal/services"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	r := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	db, err := initDatabase(cfg.DbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	defer db.Close()
	log.Info().Msg("Connected to the DB")

	adminData := data.NewAdminRepository(db)
	userData := data.NewUserRepository(db)
	blockedData := data.NewBlockedRepository(db)

	tokenManager := auth.NewTokenManager(cfg.PasetoSecret)
	adminService := services.NewAdminService(adminData)

	authService := services.NewAuthService(adminData, *tokenManager)
	userService := services.NewUserService(userData)
	blockedService := services.NewBlockedService(blockedData)

	adminHandler := handlers.NewAdminHandler(adminService)
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	blockedHandler := handlers.NewBlockedHandler(blockedService, userService)

	router := routes.NewRouter(r, adminHandler, authHandler, userHandler, blockedHandler, tokenManager)

	router.Router()

	if err := r.Run(":9001"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}

func initDatabase(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}
