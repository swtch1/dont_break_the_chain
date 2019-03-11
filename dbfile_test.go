package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const (
	contentBasic = `
last_marked_date: 2006 January 2
`
	contentChanged = `
last_marked_date: 2019 December 30
`
)

func TestLoadingYearAndYearDaysFromFile(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name         string
		fileContents string
		expectedDate string
	}{
		{
			name:         "basic",
			fileContents: contentBasic,
			expectedDate: "2006 January 2",
		},
		{
			name:         "modified",
			fileContents: contentChanged,
			expectedDate: "2019 December 30",
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
			assert.EqualValues(tt.expectedDate, dbf.date)
		})
	}
}

func TestWritingDateToDbFile(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name           string
		date           string
		expectParseErr bool
	}{
		{
			name:           "basic_write",
			date:           "2020 February 19",
			expectParseErr: false,
		},
		{
			name:           "invalid_year",
			date:           "20201 February 19",
			expectParseErr: true,
		},
		{
			name:           "invalid_Month",
			date:           "2020 Jamuary 19",
			expectParseErr: true,
		},
		{
			name:           "invalid_day",
			date:           "2020 February 35",
			expectParseErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create our test output file
			tmp, err := ioutil.TempFile("", "")
			assert.Nil(err)
			assert.Nil(tmp.Close())

			db := dbFile{path: tmp.Name()}
			err = db.WriteDate(tt.date)
			if tt.expectParseErr {
				assert.Error(err)
			} else {
				assert.Nil(err)
				assert.Nil(db.Load())
				assert.EqualValues(tt.date, db.date)
			}
		})
	}
}
