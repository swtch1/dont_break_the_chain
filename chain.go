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

// FirstMarkedDate returns the first marked date of the chain.
func (c *chain) FirstMarkedDate() (time.Time, error) {
	d, err := time.Parse(dateFmt, c.db.FirstDate())
	if err != nil {
		log.Error().Msgf("unable to parse first marked date '%s' into date format '%s'", c.db.FirstDate(), dateFmt)
		return time.Time{}, err
	}
	return d, nil
}

// LastMarkedDate returns the last marked date of the chain.
func (c *chain) LastMarkedDate() (time.Time, error) {
	d, err := time.Parse(dateFmt, c.db.LastDate())
	if err != nil {
		log.Error().Msgf("unable to parse last marked date '%s' into date format '%s'", c.db.LastDate(), dateFmt)
		return time.Time{}, err
	}
	return d, nil
}

// MarkToday marks today as done in the database, adding to the existing chain.
func (c *chain) MarkToday() error {
	t := today()
	log.Debug().Msgf("marking today '%s' in database", t)
	return c.db.WriteLastDate(t)
}

// Start will mark today as the first day in the chain.
func (c *chain) Start() error {
	t := today()
	log.Debug().Msgf("starting chain with today '%s'", t)
	return c.db.WriteFirstDate(t)
}

// Broken reports if the chain has been broken.
func (c *chain) Broken() bool {
	if c.db.LastDate() == today() || c.db.LastDate() == yesterday() {
		return false
	}
	return true
}

// Length returns the number of days the chain has been unbroken for.
func (c *chain) Length() int {
	first, err := time.Parse(dateFmt, c.db.FirstDate())
	if err != nil {
		// not sure we want a basic length func to return an error
		// we should only see this error if the date format is incorrect
		// or if the first date is wrong, which is written by this package
		log.Fatal().Err(err)
	}
	last, err := time.Parse(dateFmt, c.db.LastDate())
	if err != nil {
		log.Fatal().Err(err)
	}

	hoursDiff := last.Sub(first).Hours()

	// when starting the chain the first day and last will
	// always be the same. we want to always add 1 day to
	// the len to get an accurate reading while maintaining
	// a common sense db format
	return int(hoursDiff/24) + 1
}
