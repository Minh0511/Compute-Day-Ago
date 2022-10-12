package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
}

func computeDayAgo(endTime time.Time, limit int) string {
	if limit > 30 || limit < 1 {
		return "Error Limit Input"
	}
	var locationVN = time.FixedZone("Vietnam", 7*3600)
	endTime = endTime.In(locationVN)
	yy, mm, dd := endTime.In(locationVN).Date()
	startOfDay := time.Date(yy, mm, dd, 0, 0, 0, 0, locationVN)
	const day = 24 * time.Hour
	if startOfDay.Equal(endTime) {
		return startOfDay.Add(-time.Duration(limit) * day).UTC().String()
	}
	return startOfDay.Add(-time.Duration(limit-1) * day).UTC().String()
}

func newTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t.UTC()
}
