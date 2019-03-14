package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestChainLength(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name            string
		firstMarkedDate string
		lastMarkedDate  string
		// expected length of the chain
		expectedLen int
	}{
		{
			name:            "same_day",
			firstMarkedDate: "2000 January 1",
			lastMarkedDate:  "2000 January 1",
			expectedLen:     1,
		},
		{
			name:            "next_day",
			firstMarkedDate: "2000 January 1",
			lastMarkedDate:  "2000 January 2",
			expectedLen:     2,
		},
		{
			name:            "crossing_month_boundary", // time.Parse should handle this
			firstMarkedDate: "2000 January 31",
			lastMarkedDate:  "2000 February 1",
			expectedLen:     2,
		},
		{
			name:            "crossing_year_boundary", // time.Parse should handle this
			firstMarkedDate: "2000 December 31",
			lastMarkedDate:  "2001 January 1",
			expectedLen:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp, err := ioutil.TempFile("", "")
			assert.Nil(err)
			assert.Nil(tmp.Close())
			defer os.Remove(tmp.Name())

			db := dbFile{path: tmp.Name()}
			db.firstDate = tt.firstMarkedDate
			db.lastDate = tt.lastMarkedDate
			c := chain{db: &db}
			assert.EqualValues(tt.expectedLen, c.Length())
		})
	}
}
