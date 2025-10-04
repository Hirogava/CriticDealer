package funcgraf

import (
	"fmt"

	"github.com/Hirogava/ParkingDealer/internal/models/graf"
)

func ProcessGraphs(graphs []*models.Graph, crashPoints []models.CrashPoint, globalIndex, criticalIndex float64) ([]*models.Graph, error) {
	if len(graphs) == 0 {
		return nil, fmt.Errorf("графа нету")
	}
	if len(crashPoints) == 0 {
		return nil, fmt.Errorf("нет индекса")
	}
	crashMap := make(map[string]models.CrashPoint)
	for _, cp := range crashPoints{
		crashMap[cp.NodeID] = cp
	}
	var result []*models.Graph

    for _, g := range graphs {
        var total, count float64
		for nodeID, node := range g.Nodes{
			if cp, exists := crashMap[nodeID]; exists{
				if cp.LocalCrashIndex >= criticalIndex || cp.LocalCrashIndex > globalIndex{
					node.Danger = true

					for i := range g.Rebra{
						e := &g.Rebra[i]
						if e.From == nodeID || e.To == nodeID{
							e.Danger = true
						}
					}
				}
				total += cp.LocalCrashIndex
				count ++
			}
		}

		
    }

    return result, nil
}