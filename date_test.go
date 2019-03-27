package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGettingMonthsBetweenDates(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		err      error
		expected []Month
	}{
		{
			name:  "err_on_end_month_before_start_month",
			start: time.Date(2000, time.February, 0, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2000, time.January, 0, 0, 0, 0, 0, time.UTC),
			err:   ErrStartDateBeforeEndDate,
		},
		{
			name:  "err_on_end_year_before_start_year",
			start: time.Date(2001, time.January, 0, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2000, time.January, 0, 0, 0, 0, 0, time.UTC),
			err:   ErrStartDateBeforeEndDate,
		},
		{
			name:  "err_on_end_time_before_start_time",
			start: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			end:   time.Date(2000, time.January, 0, 0, 0, 0, 0, time.UTC),
			err:   ErrStartDateBeforeEndDate,
		},
		{
			name:     "same_month",
			start:    time.Date(2000, time.January, 5, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			expected: []Month{{2000, time.January}},
		},
		{
			name:     "next_month",
			start:    time.Date(2000, time.January, 5, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2000, time.February, 10, 0, 0, 0, 0, time.UTC),
			expected: []Month{{2000, time.January}, {2000, time.February}},
		},
		{
			name:     "several_months",
			start:    time.Date(2000, time.January, 5, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2000, time.March, 10, 0, 0, 0, 0, time.UTC),
			expected: []Month{{2000, time.January}, {2000, time.February}, {2000, time.March}},
		},
		{
			name:     "passing_year_boundary",
			start:    time.Date(2000, time.December, 5, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2001, time.February, 10, 0, 0, 0, 0, time.UTC),
			expected: []Month{{2000, time.December}, {2001, time.January}, {2001, time.February}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := monthsBetweenDates(tt.start, tt.end)
			if tt.err != nil {
				assert.Error(err)
			} else {
				assert.Nil(err)
				assert.Equal(tt.expected, m)
			}
		})
	}
}
