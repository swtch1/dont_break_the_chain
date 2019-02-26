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
