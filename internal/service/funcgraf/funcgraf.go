package funcgraf

import (
	"fmt"

	"github.com/Hirogava/ParkingDealer/internal/models/graf"
	"github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func ProcessRoutesFromAPI(routingResponse *routresponse.RouteResponse,crashPoints []graf.CrashPoint, globalIndex, criticalIndex float64,) ([]*graf.Graph, error) {
	if routingResponse.Status != "OK" {
		return nil, fmt.Errorf("ошибка API: %s", routingResponse.Status)
	}
	if len(crashPoints) == 0 {
		return nil, fmt.Errorf("нет аварийных точек")
	}

	crashMap := make(map[string]graf.CrashPoint)
	for _, cp := range crashPoints {
		crashMap[cp.NodeID] = cp
	}

	var result []*graf.Graph

	for _, route := range routingResponse.Result {

		graph := &graf.Graph{
			ID:    route.ID,
			Nodes: make(map[string]*graf.Node),
			Rebra: []graf.Rebro{},
			RiskScore: 0,
		}

		var total, count float64
		hasDanger := false

		// Создаём узлы из точек маршрута
		for i, wp := range route.Waypoints {
			nodeID := fmt.Sprintf("%s_%d", route.ID, i)

			node := &graf.Node{
				ID:      nodeID,
				Shirota: wp.ProjectedPoint.Lat,
				Dolgota: wp.ProjectedPoint.Lon,
				Danger:  false,
			}

			// Если есть данные о ДТП для этой точки
			if cp, exists := crashMap[nodeID]; exists {
				// Проверка опасности точки
				if cp.LocalCrashIndex >= criticalIndex || cp.LocalCrashIndex > globalIndex {
					node.Danger = true
					hasDanger = true
				}

				total += cp.LocalCrashIndex
				count++
			}

			graph.Nodes[nodeID] = node
		}

		// Создаём рёбра между соседними узлами маршрута
		waypointCount := len(route.Waypoints)
		for i := 0; i < waypointCount-1; i++ {
			fromID := fmt.Sprintf("%s_%d", route.ID, i)
			toID := fmt.Sprintf("%s_%d", route.ID, i+1)

			edge := graf.Rebro{
				From:   fromID,
				To:     toID,
				Danger: false,
			}

			if graph.Nodes[fromID].Danger || graph.Nodes[toID].Danger {
				edge.Danger = true
			}

			graph.Rebra = append(graph.Rebra, edge)
		}

		var riskScore float64
		if count > 0 {
			riskScore = total / count
			if riskScore > 1 {
				riskScore = 1
			}
			if riskScore < 0 {
				riskScore = 0
			}
		}

		if hasDanger {
			result = append(result, graph)
		}
		result = append(result, graph)
	}

	return result, nil
}