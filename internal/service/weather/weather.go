package weather

import (
	"fmt"
	"slices"

	routresponse "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func GetCurrentWeather(w *routresponse.WeatherResponse) []string {
	var weather []string

	fmt.Println(w,11)

	for _, wea := range w.Weather {
		if !slices.Contains(weather, wea.Main) {
			weather = append(weather, wea.Main)
		}
	}

	fmt.Println(weather, 12)

	return weather
}
