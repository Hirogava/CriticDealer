package weather

import (
	"slices"

	routresponse "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func GetCurrentWeather(w *routresponse.WeatherResponse) []string {
	var weather []string

	for _, wea := range w.Weather {
		if !slices.Contains(weather, wea.Main) {
			weather = append(weather, wea.Main)
		}
	}

	return weather
}
