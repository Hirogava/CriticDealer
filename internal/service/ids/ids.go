package ids

import (
	"hash/crc32"
	models "github.com/Hirogava/ParkingDealer/internal/models/routresponse"
)

func GetCurrentIds(r *models.RouteResponse) []int64 {
	var ids []int64

	for _, res := range r.Result {
		for _, man := range res.Maneuvers {
			hash := crc32.ChecksumIEEE([]byte(man.Comment))
			ids = append(ids, int64(hash))
		}
	}

	return ids
}
