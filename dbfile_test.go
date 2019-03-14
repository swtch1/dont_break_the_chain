package main

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

const (
	contentBasic = `
first_marked_date: 2006 January 2
last_marked_date: 2006 January 2
`
	contentChanged = `
first_marked_date: 2019 December 15
last_marked_date: 2019 December 30
`
)

func TestLoadingValuesFromDbFile(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name                    string
		fileContents            string
		expectedFirstMarkedDate string
		expectedLastMarkedDate  string
	}{
		{
			name:                    "basic",
			fileContents:            contentBasic,
			expectedFirstMarkedDate: "2006 January 2",
			expectedLastMarkedDate:  "2006 January 2",
		},
		{
			name:                    "modified",
			fileContents:            contentChanged,
			expectedFirstMarkedDate: "2019 December 15",
			expectedLastMarkedDate:  "2019 December 30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// write file contents to a temp file
			tmp, err := ioutil.TempFile("", "")
			assert.Nil(err)
			defer os.Remove(tmp.Name())
			_, err = tmp.WriteString(tt.fileContents)
			assert.Nil(err)
			assert.Nil(tmp.Close())

			// load temp file contents into new dbfile
			dbf := dbFile{path: tmp.Name()}
			assert.Nil(dbf.Load())
			assert.EqualValues(tt.expectedFirstMarkedDate, dbf.firstDate)
			assert.EqualValues(tt.expectedLastMarkedDate, dbf.lastDate)
		})
	}
}

func TestWritingLastDateToDbFile(t *testing.T) {
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
			defer os.Remove(tmp.Name())

			db := dbFile{path: tmp.Name()}
			err = db.WriteLastDate(tt.date)
			if tt.expectParseErr {
				assert.Error(err)
			} else {
				assert.Nil(err)
				assert.Nil(db.Load())
				assert.EqualValues(tt.date, db.lastDate)
			}
		})
	}
}

func TestWritingDateToExistingFilePreservesExistingValues(t *testing.T) {
	assert := assert.New(t)

	// loadedDbFile loads up and returns a dbFile
	loadedDbFile := func(path string) dbFile {
		// read vals from path into dbYaml
		y := dbYaml{}
		b, err := ioutil.ReadFile(path)
		assert.Nil(err)
		assert.Nil(yaml.Unmarshal(b, &y))

		// load vals into dbFile
		dbf := dbFile{path: path}
		assert.Nil(dbf.Load())
		return dbf
	}

	tests := []struct {
		name                   string
		initialFirstMarkedDate string
		initialLastMarkedDate  string
		updatedFirstMarkedDate string
		updatedLastMarkedDate  string
	}{
		{
			name:                   "no_changes",
			initialFirstMarkedDate: "2006 January 1",
			initialLastMarkedDate:  "2006 January 1",
			updatedFirstMarkedDate: "2010 February 5",
			updatedLastMarkedDate:  "2010 May 23",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// make a temp file for writing and reading test data
			tmp, err := ioutil.TempFile("", "")
			assert.Nil(err)
			assert.Nil(tmp.Close())
			defer os.Remove(tmp.Name())

			// write vals out to tmp file
			dbOut := dbYaml{FirstDate: tt.initialFirstMarkedDate, LastDate: tt.initialLastMarkedDate}
			y, err := yaml.Marshal(dbOut)
			assert.Nil(err)
			assert.Nil(ioutil.WriteFile(tmp.Name(), y, 0655))

			// load dbFile and ensure vals were loaded correctly
			dbf := loadedDbFile(tmp.Name())
			assert.EqualValues(tt.initialFirstMarkedDate, dbf.firstDate)
			assert.EqualValues(tt.initialLastMarkedDate, dbf.lastDate)

			// write out first date
			assert.Nil(dbf.WriteFirstDate(tt.updatedFirstMarkedDate))

			// load dbFile and ensure vals were loaded correctly
			// we expect the first value to have changed but not the last
			dbf = loadedDbFile(tmp.Name())
			assert.EqualValues(tt.updatedFirstMarkedDate, dbf.firstDate)
			assert.EqualValues(tt.initialLastMarkedDate, dbf.lastDate)

			// write out last date
			assert.Nil(dbf.WriteLastDate(tt.updatedLastMarkedDate))

			// load dbFile and ensure vals were loaded correctly
			// we expect for both values to have been changed since we changed the first above and we've
			// been using the same file the whole time.
			dbf = loadedDbFile(tmp.Name())
			assert.EqualValues(tt.updatedFirstMarkedDate, dbf.firstDate)
			assert.EqualValues(tt.updatedLastMarkedDate, dbf.lastDate)
		})
	}
}
