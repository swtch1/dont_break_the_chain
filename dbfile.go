package main

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type dbYaml struct {
	// Date matches the format '2006 January 2'
	Date string `yaml:"last_marked_date"`
}

// dbFile is the file representation of a chains database.
type dbFile struct {
	// date is the last marked day from the YAML file. this matches the format '2006 January 2'.
	date string
	// path is where the db file is located
	path string
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

	if y.Date == "" {
		log.Debug().Msgf("no date read from database file, the file may be empty")
	} else {
		log.Debug().Msgf("latest marked date in database file is %s", y.Date)
	}
	db.date = y.Date
	return nil
}

// WriteDate writes the date string to the dbFile. The format
// of the date should be YYYY time.Month DD.  An invalid date will
// result in an error.
//
// Example: db.WriteDate("2020 February 5")
func (db *dbFile) WriteDate(date string) error {
	// validate date before writing
	if _, err := time.Parse(dateFmt, date); err != nil {
		return err
	}

	y, err := yaml.Marshal(dbYaml{Date: date})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(db.path, y, 0655)
}

func (db *dbFile) Date() string {
	return db.date
}
