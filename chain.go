package main

import (
	"github.com/rs/zerolog/log"
	"time"
)

type DaterWriter interface {
	// FirstDate is the string date (format YYYY time.Month DD) from the first marked day of the chain.
	FirstDate() string
	// LastDate is the string date (format YYYY time.Month DD) from the last marked day of the chain.
	LastDate() string
	// WriteFirstDate writes date to the first marked day of the chain in the database.
	WriteFirstDate(date string) error
	// WriteLastDate writes date to the last marked day of the chain in the database.
	WriteLastDate(date string) error
}

type chain struct {
	// chain database
	db DaterWriter
}

// MarkToday marks today as done in the database.
func (c *chain) MarkToday() error {
	t := today()
	log.Debug().Msgf("marking today '%s' in database", t)
	return c.db.WriteLastDate(t)
}

// Broken reports if the chain has been broken.
func (c *chain) Broken() bool {
	if c.db.LastDate() == today() || c.db.LastDate() == yesterday() {
		return false
	}
	return true
}

// Length returns the number of days the chain has been unbroken for.
func (c *chain) Length() (int, error) {
	d, err := time.Parse(dateFmt, c.db.FirstDate()) // FIXME
	if err != nil {
		return 0, err
	}
	hoursDiff := time.Now().Sub(d).Hours()
	return int(hoursDiff / 24), nil
}
