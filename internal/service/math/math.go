package math

import (
	"hash/crc32"
	
	dbModel "github.com/Hirogava/ParkingDealer/internal/models/db"
	models "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func CountCurrentCriticality(r *models.RouteResponse, criticals map[int64][]dbModel.Critical, weather []string) {
	for _, route := range r.Result {
		for i := range route.Maneuvers {
			m := &route.Maneuvers[i]

			length := float64(m.OutcomingPath.Distance)
			if length == 0 {
				length = 1
			}

			hash := int64(crc32.ChecksumIEEE([]byte(m.Comment)))
			accidents := criticals[hash]

			var (
				totalAccidents          float64
				clearAccidents          float64
				currentWeatherAccidents float64
				totalTraffic            float64
			)

			for _, acc := range accidents {
				totalAccidents++
				totalTraffic += float64(acc.Traffic)

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

			avgTraffic := totalTraffic / totalAccidents
			if avgTraffic <= 0 {
				avgTraffic = 1
			}

			clearRatio := clearAccidents / totalAccidents

			var probability float64
			if currentWeatherAccidents == clearAccidents || currentWeatherAccidents == 0 {
				probability = clearAccidents / (length * avgTraffic)
			} else {
				probability = (currentWeatherAccidents / (length * avgTraffic)) +
					(totalAccidents / (length * avgTraffic) * clearRatio)
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

func F32(v float32) *float32 { return &v }
