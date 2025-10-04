package funcgraf

import (
	"fmt"
	"github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func ProcessRoutesFromAPI(routingResponse *routresponse.RouteResponse, globalIndex float64) (*routresponse.RouteResponse, error) {
	if routingResponse == nil {
		return nil, fmt.Errorf("пустой ответ API")
	}

	if routingResponse.Status != "OK" {
		return nil, fmt.Errorf("ошибка API: %s", routingResponse.Status)
	}

	if len(routingResponse.Result) == 0 {
		return nil, fmt.Errorf("в ответе нет маршрутов")
	}
	for _, route := range routingResponse.Result {
		for i := range route.Maneuvers {
			maneuver := &route.Maneuvers[i]

			localRisk := 0.0
			if maneuver.CriticalProbability != nil {
				localRisk = float64(*maneuver.CriticalProbability)
			}

			// Определяем опасность
			if (maneuver.Critical != nil && *maneuver.Critical) || localRisk > globalIndex {
				maneuver.Critical = newBool(true) // явная отметка опасности
			} else {
				maneuver.Critical = newBool(false)
			}
		}
	}

	return routingResponse, nil

}

func newBool(b bool) *bool {
	return &b
}

