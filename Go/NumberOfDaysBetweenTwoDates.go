package Go

import (
	"strconv"
	"time"
)

func DaysBetweenDates(date1 string, date2 string) int {
	t1y, _ := strconv.Atoi(date1[:4])
	t1m, _ := strconv.Atoi(date1[5:7])
	t1d, _ := strconv.Atoi(date1[8:])
	t1 := date(t1y, t1m, t1d)
	t2y, _ := strconv.Atoi(date2[:4])
	t2m, _ := strconv.Atoi(date2[5:7])
	t2d, _ := strconv.Atoi(date2[8:])
	t2 := date(t2y, t2m, t2d)

	return abs(int(t2.Sub(t1).Hours() / 24))
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}
