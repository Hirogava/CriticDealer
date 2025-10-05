package http

import (
	"time"

	"github.com/Hirogava/ParkingDealer/internal/handler/api"
	mgr "github.com/Hirogava/ParkingDealer/internal/repository/postgres/api"
	"github.com/gin-contrib/cors"

	"github.com/Hirogava/ParkingDealer/internal/config/logger"

	"github.com/gin-gonic/gin"
)

func CreateRouter(manager *mgr.Manager) *gin.Engine {
	logger.Logger.Debug("Creating HTTP router")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api.InitParams(r, manager)

	logger.Logger.Info("HTTP router created successfully")
	return r
}
