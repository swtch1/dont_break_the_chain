package main

import (
	"github.com/rs/zerolog/log"
	"time"
)

type DaterWriter interface {
	// Date is the string date (format YYYY time.Month DD) from the last marked day.
	Date() string
	WriteDate(date string) error
}

type chain struct {
	// chain database
	db DaterWriter
}

// markToday marks today as done in the database.
func (c *chain) markToday() error {
	t := today()
	log.Debug().Msgf("marking today '%s' in database", t)
	return c.db.WriteDate(t)
}

// Broken reports if the chain has been broken.
func (c *chain) Broken() bool {
	if c.db.Date() == today() || c.db.Date() == yesterday() {
		return false
	}
	return true
}

// Length returns the number of days the chain has been unbroken for.
func (c *chain) Length() (int, error) {
	d, err := time.Parse(dateFmt, c.db.Date())
	if err != nil {
		return 0, err
	}
	hoursDiff := time.Now().Sub(d).Hours()
	return int(hoursDiff / 24), nil
}
