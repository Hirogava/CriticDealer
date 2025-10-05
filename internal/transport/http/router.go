package http

import (
	"github.com/Hirogava/ParkingDealer/internal/handler/api"
	mgr "github.com/Hirogava/ParkingDealer/internal/repository/postgres/api"

	"github.com/Hirogava/ParkingDealer/internal/config/logger"

	"github.com/gin-gonic/gin"
)

func CreateRouter(manager *mgr.Manager) *gin.Engine {
	logger.Logger.Debug("Creating HTTP router")

	r := gin.Default()

	api.InitParams(r, manager)

	logger.Logger.Info("HTTP router created successfully")
	return r
}
