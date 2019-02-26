package main

import (
	"os"
)

type dbYaml struct {
	YearDays int `yaml:"last_marked_year_days"`
	Year     int `yaml:"last_marked_year"`
}

// dbFile is the file representation of a chains database.
type dbFile struct {
	path string
}

func (db dbFile) Write(p []byte) (n int, err error) {
	f, err := os.Create(db.path)
	if err != nil {
		return 0, err
	}

	return f.Write(p)
}

func (db dbFile) WriteDate(yearDays, year int) error {
	return nil
}

func (db dbFile) Days() int {
	return 12345
}

func (db dbFile) Year() int {
	return 2000
}
