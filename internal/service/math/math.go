package math

import (
	models "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
	dbModel "github.com/Hirogava/ParkingDealer/internal/models/db"
)


func CountCurrentCriticality(r *models.RouteResponse, criticals []dbModel.Critical, globalKoef float64) {}