package math

import (
	"strconv"

	dbModel "github.com/Hirogava/ParkingDealer/internal/models/db"
	models "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func CountCurrentCriticality(r *models.RouteResponse, criticals map[int][]dbModel.Critical, weather []string) {
	for _, route := range r.Result {
		for i := range route.Maneuvers {
			m := &route.Maneuvers[i]

			length := float64(m.OutcomingPath.Distance)
			if length == 0 {
				length = 1
			}

			accidents := criticals[routeIndexToMovementID(route.ID)]
			var totalAccidents, clearAccidents, currentWeatherAccidents float64

			for _, acc := range accidents {
				totalAccidents++
				for _, w := range acc.Weather {
					if w == "Clear" {
						clearAccidents++
					}
					if matchesCurrentWeather(w, weather) {
						currentWeatherAccidents++
					}
				}
			}

			if totalAccidents == 0 {
				m.CriticalProbability = F32(0)
				continue
			}

			clearRatio := clearAccidents / totalAccidents

			var probability float64
			if currentWeatherAccidents == clearAccidents || currentWeatherAccidents == 0 {
				probability = clearAccidents / length
			} else {
				probability = (currentWeatherAccidents / length) +
					(totalAccidents / length * clearRatio)
			}

			m.CriticalProbability = F32(float32(probability))
		}
	}
}

func matchesCurrentWeather(w string, currentWeather []string) bool {
	for _, cw := range currentWeather {
		if cw == w {
			return true
		}
	}
	return false
}

func routeIndexToMovementID(routeID string) int {
	intID, _ := strconv.Atoi(routeID)
	return intID
}

func F32(v float32) *float32 {
	return &v
}
