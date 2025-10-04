package api

import (
	dbModels "github.com/Hirogava/ParkingDealer/internal/models/db"
	models "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
	"github.com/Hirogava/ParkingDealer/internal/service/funcmonth"
	"github.com/Hirogava/ParkingDealer/internal/service/ids"
	"github.com/Hirogava/ParkingDealer/internal/service/math"
	"github.com/Hirogava/ParkingDealer/internal/service/weather"
	"github.com/lib/pq"
)

func (manager *Manager) GetCriticalManeuvers(r *models.RouteResponse, w *models.WeatherResponse) (*models.RouteResponse, float64, error) {
	weather := weather.GetCurrentWeather(w)
	ids := ids.GetCurrentIds(r)
	dayType, month := funcmonth.WeekdayOrWeekend()

	rows, err := manager.Conn.Query(`
		SELECT a.hash, ARRAY_AGG(w.weather_type), a.traffic
		FROM accident a
		JOIN weather w ON a.weather_id = w.id
		WHERE a.hash = ANY($1)
			AND a.day_type = $2
			AND (w.weather_type = ANY($3) OR w.weather_type = 'Clear')
			AND a.month = ANY($5)
			AND a.dtp_time BETWEEN GREATEST($4 - 1, 1) AND LEAST($4 + 1, 24)
		GROUP BY a.hash;
	`, pq.Array(ids), dayType, pq.Array(weather), funcmonth.GetCurrentHour(), pq.Array(month))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cur []dbModels.Critical
	for rows.Next() {
		var c dbModels.Critical

		if err := rows.Scan(&c.ID, pq.Array(&c.Weather), &c.Traffic); err != nil {
			return nil, 0, err
		}

		cur = append(cur, c)
	}

	criticals := make(map[int64][]dbModels.Critical)
	for _, c := range cur {
		criticals[c.ID] = append(criticals[c.ID], c)
	}

	var globalKoef float64
	if err := manager.Conn.QueryRow(`SELECT dtp_koef FROM global_accident_statistic LIMIT 1`).Scan(&globalKoef); err != nil {
		return nil, 0, err
	}

	math.CountCurrentCriticality(r, criticals, weather)
	return r, globalKoef, nil
}

