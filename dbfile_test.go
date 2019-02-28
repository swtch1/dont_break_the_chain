package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const contentBasic = `
last_marked_year_days: 65
last_marked_year: 2002
`

func TestLoadingYearAndYearDaysFromFile(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name             string
		fileContents     string
		expectedYearDays int
		expectedYear     int
	}{
		{
			name:             "year_days_set",
			fileContents:     contentBasic,
			expectedYearDays: 65,
			expectedYear:     2002,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// write file contents to a temp file
			f, err := ioutil.TempFile("", "")
			assert.Nil(err)
			defer os.Remove(f.Name())
			_, err = f.WriteString(tt.fileContents)
			assert.Nil(err)
			assert.Nil(f.Close())

			// load temp file contents into new dbfile
			dbf := dbFile{path: f.Name()}
			assert.Nil(dbf.Load())
			assert.Equal(tt.expectedYearDays, dbf.yearDays)
			assert.Equal(tt.expectedYear, dbf.year)
		})
	}
}

func TestWritingDateToDbFile(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		yearDays int
		year     int
	}{
		{
			name:     "basic_write",
			yearDays: 54,
			year:     2004,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create our test output file
			tmp, err := ioutil.TempFile("", "")
			assert.Nil(err)
			assert.Nil(tmp.Close())

			db := dbFile{path: tmp.Name()}
			assert.Nil(db.WriteDate(tt.yearDays, tt.year))
			assert.Nil(db.Load())

			assert.EqualValues(tt.yearDays, db.yearDays)
			assert.EqualValues(tt.year, db.year)
		})
	}
}
