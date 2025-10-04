package funcmonth

import (
	"time"

)

func WeekdayOrWeekend() (string, int) {
	now := time.Now()
	weekday := now.Weekday()
	month := int(now.Month())
	if weekday == time.Saturday || weekday == time.Sunday {
		return "weekend", month
	}
	return "weekday", month
}

