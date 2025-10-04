package funcgraf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Hirogava/ParkingDealer/internal/models/graf"
	"github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func ProcessRoutesFromAPI(routingResponse *routresponse.RouteResponse, globalIndex float64) ([]*graf.Graph, error) {
	if routingResponse == nil {
		return nil, fmt.Errorf("пустой ответ API")
	}

	if routingResponse.Status != "OK" {
		return nil, fmt.Errorf("ошибка API: %s", routingResponse.Status)
	}

	if len(routingResponse.Result) == 0 {
		return nil, fmt.Errorf("в ответе нет маршрутов")
	}

	var result []*graf.Graph

	for _, route := range routingResponse.Result {
		graph := &graf.Graph{
			ID:    route.ID,
			Nodes: make(map[string]*graf.Node),
			Rebra: []graf.Rebro{},
		}

		var totalRisk float64
		var riskCount int

		// Обрабатываем манёвры
		for i, maneuver := range route.Maneuvers {
			nodeID := fmt.Sprintf("%s_%d", route.ID, i)

			var lat, lon float64
			if len(maneuver.OutcomingPath.Geometry) > 0 {
				lat, lon = parseLatLon(maneuver.OutcomingPath.Geometry[0].Selection)
			}
			node := &graf.Node{
				ID:      nodeID,
				Shirota: lat,
				Dolgota: lon,
				Danger:  false,
			}

			localRisk := 0.0
			if maneuver.CriticalProbability != nil {
				localRisk = float64(*maneuver.CriticalProbability)
			}

			if (maneuver.Critical != nil && *maneuver.Critical) || localRisk > globalIndex {
				node.Danger = true
			}

			graph.Nodes[nodeID] = node
			totalRisk += localRisk
			riskCount++
		}

		// Формируем рёбра между последовательными манёврами
		for i := 0; i < len(route.Maneuvers)-1; i++ {
			fromID := fmt.Sprintf("%s_%d", route.ID, i)
			toID := fmt.Sprintf("%s_%d", route.ID, i+1)

			rebro := graf.Rebro{
				From:   fromID,
				To:     toID,
				Danger: graph.Nodes[fromID].Danger || graph.Nodes[toID].Danger,
			}
			graph.Rebra = append(graph.Rebra, rebro)
		}

		// Рассчитываем средний риск
		if riskCount > 0 {
			graph.RiskScore = totalRisk / float64(riskCount)
			if graph.RiskScore > 1 {
				graph.RiskScore = 1
			}
		}

		result = append(result, graph)
	}

	return result, nil
}

func parseLatLon(selection string) (float64, float64) {
	parts := strings.Split(selection, ",")
	if len(parts) != 2 {
		return 0, 0
	}
	lat, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	lon, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err1 != nil || err2 != nil {
		return 0, 0
	}
	return lat, lon
}
