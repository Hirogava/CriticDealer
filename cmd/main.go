package main

import (
	"os"

	"github.com/Hirogava/ParkingDealer/internal/config/environment"
	manager "github.com/Hirogava/ParkingDealer/internal/repository/postgres/api"
	"github.com/Hirogava/ParkingDealer/internal/config/logger"
	router "github.com/Hirogava/ParkingDealer/internal/transport/http"
)

func main() {
	environment.LoadEnvFile(".env")

	logger.LogInit()
	logger.Logger.Info("Starting ParkingDealer backend server")

	dbConnStr := os.Getenv("DB_CONNECT_STRING")
	if dbConnStr == "" {
		logger.Logger.Fatal("DB_CONNECT_STRING environment variable is required")
	}
	logger.Logger.Info("Connecting to database", "connection_string", dbConnStr)

	manager := manager.NewApiManager("postgres", dbConnStr)
	logger.Logger.Info("Database connection established successfully")

	logger.Logger.Info("Initializing HTTP router")
	r := router.CreateRouter(manager)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080"
		logger.Logger.Warn("SERVER_PORT not set, using default port 8080")
	}

	logger.Logger.Info("Starting HTTP server", "port", serverPort)
	if err := r.Run(serverPort); err != nil {
		logger.Logger.Fatal("Failed to start HTTP server", "error", err)
	}
}
