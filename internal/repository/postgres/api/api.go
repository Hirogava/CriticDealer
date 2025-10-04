package api

import (
	dbModels "github.com/Hirogava/ParkingDealer/internal/models/db"
	models "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
	"github.com/Hirogava/ParkingDealer/internal/service/ids"
	"github.com/Hirogava/ParkingDealer/internal/service/weather"
	"github.com/lib/pq"
)

func (manager *Manager) GetCriticalManeuvers(r *models.RouteResponse, w *models.WeatherResponse) (*models.RouteResponse, float64, error) {
	weather := weather.GetCurrentWeather(w)
	ids := ids.GetCurrentIds(r)

	rows, err := db.Query(ctx, `
		SELECT a.id, ARRAY_AGG(w.weather_type) AS weather
		FROM accident a
		JOIN weather w ON a.weather_id = w.id
		WHERE a.movement_id = ANY($1)
			AND a.day_type = $2
			AND (w.weather_type = ANY($3) OR w.weather_type = 'Clear')
			AND a.month = ANY($5)
			AND a.dtp_time BETWEEN GREATEST($4 - 1, 1) AND LEAST($4 + 1, 24)
		GROUP BY a.id;
		`, pq.Array(ids), dayType, pq.Array(weather), currentHour, pq.Array(months))
}
