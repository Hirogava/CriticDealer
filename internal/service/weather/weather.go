package weather

import (
	"slices"

	"github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func GetCurrentWeather(w *routresponse.WeatherResponse) []string {
	var weather []string

	for _, wea := range w.Current.Weather {
		for _, main := range wea.Main {
			if !slices.Contains(weather, main) {
				weather = append(weather, main)
			}
		}
	}

	return weather
}
