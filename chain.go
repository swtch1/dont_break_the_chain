package main

import (
	"fmt"
	"io"
)

type chain struct {
	// path to database
	dbPath io.ReadWriter
}

// markToday marks today as done in the database.
func (c *chain) markToday() error {
	t := today()
	b := []byte(fmt.Sprintf("%d:%d", t.yearDays, t.year))
	if _, err := c.dbPath.Write(b); err != nil {
		return err
	}
	return nil
}

// broken reports if the chain has been broken.
func (c *chain) broken() bool {
	return true
}
