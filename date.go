package main

import (
	"fmt"
	"time"
)

const dateFmt = "2006 January 2"

// today returns the current date at a string in the format of 'YYYY time.Month DD'.
func today() string {
	now := time.Now()
	y, m, d := now.Date()
	return dateString(y, m, d)
}

// yesterday returns the date of yesterday,
func yesterday() string {
	now := time.Now()
	// subtract 24 hours to get yesterday
	oneDayAgo := now.Add(-time.Hour * 24)
	y, m, d := oneDayAgo.Date()
	return dateString(y, m, d)
}

func dateString(year int, month time.Month, day int) string {
	return fmt.Sprintf("%d %s %d", year, month, day)
}
