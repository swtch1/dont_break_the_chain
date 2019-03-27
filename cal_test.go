package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGettingCenterPosition(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name            string
		wordLen         int
		workingSpaceLen int
		// expected result
		expected int
	}{
		{"word_matches_available_space", 5, 5, 0},
		{"one_space_on_each_side", 2, 4, 1},
		{"equal_space_on_each_side", 2, 6, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.expected, CenterPos(tt.wordLen, tt.workingSpaceLen))
		})
	}
}

func TestCenterString(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name            string
		input           string
		workingSpaceLen int
		expected        string
	}{
		{"word_len_equal_to_working_space", "red", 3, "red"},
		{"word_len_greater_than_working_space", "red", 2, "red"},
		{"len_2_spaces_longer_than_word", "red", 5, " red "},
		{"len_4_spaces_longer_than_word", "red", 7, "  red  "},
		{"len_3_spaces_longer_than_word", "red", 6, " red  "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.expected, CenterString(tt.input, tt.workingSpaceLen))
		})
	}
}

func TestDateIsInRange(t *testing.T) {
	assert := assert.New(t)

	randomDate := time.Date(2000, time.March, 4, 13, 54, 16, 92, time.Local)
	day := time.Hour * 24
	year := day * 365
	tests := []struct {
		name     string
		day      time.Time
		start    time.Time
		end      time.Time
		expected bool
	}{
		{"same_time_as_start", randomDate, randomDate, randomDate.Add(day), true},
		{"same_time_as_end", randomDate, randomDate.Add(-day), randomDate, true},
		{"same_time_as_start_and_end", randomDate, randomDate, randomDate, true},

		{"date_within_range", randomDate, randomDate.Add(-day), randomDate.Add(day), true},
		{"date_before_range", randomDate, randomDate.Add(day), randomDate.Add(day * 2), false},
		{"date_after_range", randomDate, randomDate.Add(-day * 2), randomDate.Add(-day), false},

		{"same_day_later_start_year", randomDate, randomDate.Add(year), randomDate, false},
		{"same_day_earlier_end_year", randomDate, randomDate, randomDate.Add(-year), false},
		{"same_day_earlier_start_year", randomDate, randomDate.Add(-year), randomDate, true},

		// we only want to track granularity at the day level
		{"start_one_minute_behind", randomDate, randomDate.Add(time.Minute), randomDate, true},
		{"end_one_minute_behind", randomDate, randomDate, randomDate.Add(-time.Minute), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(dateInRange(tt.day, tt.start, tt.end), tt.expected)
		})
	}
}
