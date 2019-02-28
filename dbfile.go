package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type dbYaml struct {
	YearDays int `yaml:"last_marked_year_days"`
	Year     int `yaml:"last_marked_year"`
}

// dbFile is the file representation of a chains database.
type dbFile struct {
	yearDays int
	year     int
	path     string
}

// Load loads the yearDays and year from the source file.
func (db *dbFile) Load() error {
	b, err := ioutil.ReadFile(db.path)
	if err != nil {
		return err
	}

	y := dbYaml{}
	if err = yaml.Unmarshal(b, &y); err != nil {
		return err
	}

	db.yearDays = y.YearDays
	db.year = y.Year
	return nil
}

func (db *dbFile) Write(p []byte) (n int, err error) {
	f, err := os.Create(db.path)
	if err != nil {
		return 0, err
	}

	return f.Write(p)
}

func (db *dbFile) WriteDate(yearDays, year int) error {
	y, err := yaml.Marshal(dbYaml{YearDays: yearDays, Year: year})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(db.path, y, 0655)
}

func (db *dbFile) Days() int {
	return db.yearDays
}

func (db *dbFile) Year() int {
	return db.year
}
