package http

import (
	mgr "github.com/Hirogava/ParkingDealer/internal/repository/postgres/api"

	"github.com/Hirogava/ParkingDealer/internal/config/logger"

	"github.com/gin-gonic/gin"
)

func CreateRouter(manager *mgr.Manager) *gin.Engine {
	logger.Logger.Debug("Creating HTTP router")

	r := gin.Default()

	// тут инициализируем роуты
	// auth.InitAuthHandlers(r, manager) 
	// manager - работа с БД

	logger.Logger.Info("HTTP router created successfully")
	return r
}
