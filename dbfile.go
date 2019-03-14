package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type dbYaml struct {
	// FirstDate is the first parked date of the chain, matching the format '2006 January 2'.
	FirstDate string `yaml:"first_marked_date"`
	// LastDate is the last marked date of the chain, matching the format '2006 January 2'.
	LastDate string `yaml:"last_marked_date"`
}

// dbFile is the file representation of a chains database.
type dbFile struct {
	// path is where the db file is located
	path string
	// firstDate is the first marked day from the YAML file. this matches the format '2006 January 2'.
	firstDate string
	// lastDate is the last marked day from the YAML file. this matches the format '2006 January 2'.
	lastDate string
}

// Load loads the yearDays and year from the source file.
func (db *dbFile) Load() error {
	log.Debug().Msgf("loading config from database file at %s", db.path)
	b, err := ioutil.ReadFile(db.path)
	if err != nil {
		return err
	}

	y := dbYaml{}
	if err = yaml.Unmarshal(b, &y); err != nil {
		return err
	}

	if y.LastDate == "" {
		log.Debug().Msgf("no date read from database file, the file may be empty")
	} else {
		log.Debug().Msgf("latest marked date in database file is %s", y.LastDate)
	}
	db.firstDate, db.lastDate = y.FirstDate, y.LastDate
	return nil
}

// FirstDate returns the first marked date of the chain.
func (db *dbFile) FirstDate() string {
	return db.firstDate
}

// LastDate returns the last marked date of the chain.
func (db *dbFile) LastDate() string {
	return db.lastDate
}

// WriteFirstDate writes the date string to the dbFile.  The format
// of the date should be YYYY time.Month DD.  An invalid date will
// result in an error.
//
// Example: db.WriteFirstDate("2020 February 5")
func (db *dbFile) WriteFirstDate(date string) error {
	if err := validateDate(date); err != nil {
		return err
	}

	if err := db.Load(); err != nil {
		return errors.Wrapf(err, "unable to load db when attempting to write first date")
	}

	db.firstDate = date
	y, err := yaml.Marshal(dbYaml{FirstDate: db.firstDate, LastDate: db.lastDate})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(db.path, y, 0655) // FIXME: this is almost definitely going to fail because we are not parsing the file input first.
}

// WriteLastDate writes the date string to the dbFile. The format
// of the date should be YYYY time.Month DD.  An invalid date will
// result in an error.
//
// Example: db.WriteLastDate("2020 February 5")
func (db *dbFile) WriteLastDate(date string) error {
	if err := validateDate(date); err != nil {
		return err
	}

	if err := db.Load(); err != nil {
		return errors.Wrapf(err, "unable to load db when attempting to write last date")
	}

	db.lastDate = date
	y, err := yaml.Marshal(dbYaml{FirstDate: db.firstDate, LastDate: db.lastDate})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(db.path, y, 0655)
}

// validateDate ensures that the date can be parsed.
func validateDate(date string) error {
	if _, err := time.Parse(dateFmt, date); err != nil {
		return err
	}
	return nil
}
