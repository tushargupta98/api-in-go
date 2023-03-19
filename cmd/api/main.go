// @title API-In-Go
// @description This API is a template to create APIs' in Golang
// @BasePath /api/v1
// @version 1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
package main

import (
	_ "github.com/swaggo/swag"
	"github.com/tushargupta98/api-in-go/cache"
	"github.com/tushargupta98/api-in-go/config"
	"github.com/tushargupta98/api-in-go/db"
	logger "github.com/tushargupta98/api-in-go/logger"
	"github.com/tushargupta98/api-in-go/pkg/server"
)

func main() {
	cfg := config.GetConfig()

	cache := cache.NewRedisClient()
	s := server.NewServer()

	// Setup the middleware.
	s.SetupMiddleware()

	// Setup the routes.
	s.SetupRoutes(db.DB, *cache, *cfg)

	// Start the API.
	// addr := cfg.Server.Host + ":" + cfg.Server.Port
	addr := ":" + cfg.Server.Port
	err := s.ListenAndServe(addr)
	if err != nil {
		logger.Logger.Error("Error starting server:", err)
	} else {
		logger.Logger.Info("GoLang API Launch Successful.")
	}
}
