package main

import (
	"os"

	"github.com/Hirogava/ParkingDealer/internal/config/environment"
	manager "github.com/Hirogava/ParkingDealer/internal/repository/postgres/api"
	migrations "github.com/Hirogava/ParkingDealer/internal/repository/postgres"
	"github.com/Hirogava/ParkingDealer/internal/config/logger"
	router "github.com/Hirogava/ParkingDealer/internal/transport/http"
	_ "github.com/Hirogava/ParkingDealer/docs"
	swaggerFiles "github.com/swaggo/files"    
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @title SafeRoute API
// @version 1.0
// @description SafeRoute API — сервис для построения безопасных маршрутов
// @host localhost:8080
// @BasePath /
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

	logger.Logger.Info("Running database migrations")
	migrations.Migrate(manager.Conn, "api")
	logger.Logger.Info("Database migrations completed successfully")

	logger.Logger.Info("Initializing HTTP router")
	r := router.CreateRouter(manager)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080"
		logger.Logger.Warn("SERVER_PORT not set, using default port 8080")
	}

	r.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Logger.Info("Starting HTTP server", "port", serverPort)
	if err := r.Run(serverPort); err != nil {
		logger.Logger.Fatal("Failed to start HTTP server", "error", err)
	}
}
