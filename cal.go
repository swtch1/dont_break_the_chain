package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// PrintXCal prints a calendar of month to stdout, marking
// start through end dates with 'X'.  Use time.Unix(0, 0) for start
// and end to exclude 'X' marking dates.
func PrintXCal(year int, month time.Month, start, end time.Time) {
	ng := NewNumGen()

	// return at end of calendar
	defer fmt.Println()

	// get month and year to print above calendar
	monthAndYear := fmt.Sprintf("%s %d", month.String(), year)

	// print month and year, using length of days of the week header to center above the calendar
	dayOfWeekHeader := "Su Mo Tu We Th Fr Sa"
	fmt.Println(CenterString(monthAndYear, len(dayOfWeekHeader)))

	fmt.Println(dayOfWeekHeader)

	// check the weekday the first falls on so we know how many spaces to skip
	weekdayIndex := int(firstWeekdayOfMonth(year, month))
	for i := 0; i < weekdayIndex; i++ {
		fmt.Print("   ")
	}

	// dayIndex tracks which day we're on so we know when to newline, taking the start week day into account
	dayIndex := weekdayIndex

	// print the calendar
	var dayNum int
	daysThisMonth := daysInMonth(year, month)
	for dayNum < daysThisMonth {
		day := ng.Next()
		dayNum, _ = strconv.Atoi(strings.TrimLeft(day, " "))

		date := time.Date(year, month, dayNum, 0, 0, 0, 0, time.Local)
		if dateInRange(date, start, end) {
			day = " X"
		}

		fmt.Print(day, " ") // spacing between numbers

		dayIndex += 1
		if dayIndex == 7 && dayNum != daysThisMonth {
			fmt.Println()
			dayIndex = 0
		}
	}
}

// CenterPos returns the position number at which a word, of length workLen, should start
// to be centered in the middle of a space of size workingSpaceLen.
//
// For example, in a space of size 5, the word 'red' would have one empty space in front
// and one space behind to be centered in the middle of the available space.
func CenterPos(wordLen int, workingSpaceLen int) int {
	if workingSpaceLen <= wordLen {
		return 0
	}

	return (workingSpaceLen - wordLen) / 2
}

// CenterString returns the given string, s, centered in the workingSpaceLen.
func CenterString(s string, workingSpaceLen int) string {
	ls := len(s)
	if ls >= workingSpaceLen {
		return s
	}

	var r string
	var n int // number of chars written
	for i := 0; i < CenterPos(ls, workingSpaceLen); i++ {
		n += 1
		r += " "
	}

	// add length of passed string to tracker
	n += len(s)

	// add passed string to centered return string
	r += s

	// fill the rest of the working len with spaces
	for n < workingSpaceLen {
		n += 1
		r += " "
	}
	return r
}

// daysInMonth returns the number of days in the given year and month.
func daysInMonth(year int, month time.Month) int {
	// date of first day in year and month
	d := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// add a month
	d = d.AddDate(0, 1, 0)
	// subtract a day
	d = d.AddDate(0, 0, -1)
	return d.Day()
}

// firstWeekdayOfMonth returns the time.Weekday that the 1st of the month falls on.
func firstWeekdayOfMonth(year int, month time.Month) time.Weekday {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday()
}

// dateInRange asserts that the day is one of, or between, start and end. The granularity
// is only tracked at a day level, so off by seconds, minutes or hours will still return
// true if the date is on the same day as start or end.
func dateInRange(day, start, end time.Time) bool {
	// only track day granularity
	day = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())

	if day.Sub(start) < 0 || day.Sub(end) > 0 {
		return false
	}
	return true
}
