package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	currentDay := current.Format("2006-01-02")
	yesterday := current.AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayZero, _ := time.ParseInLocation("2006-01-02", yesterday, time.Local)
	yesterdayLast, _ := time.ParseInLocation("2006-01-02 15:04:05", yesterday+" 23:59:59", time.Local)
	fmt.Print(currentDay, yesterday, yesterdayZero.Unix(), yesterdayLast.Unix())
}
