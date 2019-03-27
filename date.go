package main

import (
	"fmt"
	"time"
)

const dateFmt = "2006 January 2"

// today returns the current date, formatted with dateString.
func today() string {
	now := time.Now()
	y, m, d := now.Date()
	return dateString(y, m, d)
}

// yesterday returns the date of yesterday, formatted with dateString.
func yesterday() string {
	now := time.Now()
	// subtract 24 hours to get yesterday
	oneDayAgo := now.Add(-time.Hour * 24)
	y, m, d := oneDayAgo.Date()
	return dateString(y, m, d)
}

// dateString will convert year month and day into a string in the format of 'YYYY time.Month DD'.
func dateString(year int, month time.Month, day int) string {
	return fmt.Sprintf("%d %s %d", year, month, day)
}

// Month holds a month and associated year.
type Month struct {
	Year  int
	Month time.Month
}

// monthsBetweenDates finds all months including and in between start and end.
func monthsBetweenDates(start, end time.Time) ([]Month, error) {
	var months []Month

	if start.Sub(end) > 0 {
		return []Month{}, ErrStartDateBeforeEndDate
	}

	// track and add to the current month as we move toward end
	var d = start
	//var m = Month{start.Year(), start.Month()}
	for {
		if d.Month() == end.Month() && d.Year() == end.Year() {
			break
		}

		months = append(months, Month{d.Year(), d.Month()})
		d = d.AddDate(0, 1, 0)
	}
	months = append(months, Month{Year: end.Year(), Month: end.Month()})
	return months, nil
}
