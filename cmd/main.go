package main

import (
	"os"

	"github.com/Hirogava/ParkingDealer/internal/config/environment"
	"github.com/Hirogava/ParkingDealer/internal/config/logger"
	router "github.com/Hirogava/ParkingDealer/internal/transport/http"
)

func main() {
	environment.LoadEnvFile(".env")

	logger.LogInit()
	logger.Logger.Info("Starting ParkingDealer backend server")

	logger.Logger.Info("Initializing HTTP router")
	r := router.CreateRouter()

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
