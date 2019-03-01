package main

import "time"

type date struct {
	yearDays int
	year     int
}

// today returns the current date.
func today() date {
	now := time.Now()
	return date{
		yearDays: now.YearDay(),
		year:     now.Year(),
	}
}

// yesterday returns the date of yesterday,
func yesterday() date {
	now := time.Now()
	// subtract 24 hours to get yesterday
	oneDayAgo := now.Add(-time.Hour * 24)
	return date{
		yearDays: oneDayAgo.YearDay(),
		year:     oneDayAgo.Year(),
	}
}
