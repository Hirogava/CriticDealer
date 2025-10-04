package funcmonth

import "time"

func GetCurrentHour() int {
	return time.Now().Hour()
}