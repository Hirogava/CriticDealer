package ids

import (
	"strconv"

	models "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func GetCurrentIds(r *models.RouteResponse) []int {
	var ids []int

	for _, res := range r.Result {
		for _, man := range res.Maneuvers {
			intId, _ := strconv.Atoi(man.ID)
			ids = append(ids, intId)
		}
	}

	return ids
}