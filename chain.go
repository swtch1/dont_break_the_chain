package main

type DaterWriter interface {
	Days() int // last recorded numbered day of the year is in time.Now().YearDay()
	Year() int // last recorded year as in time.Now().Year()
	WriteDate(yearDays, year int) error
}

type chain struct {
	// chain database
	db DaterWriter
}

// markToday marks today as done in the database.
func (c *chain) markToday() error {
	return c.db.WriteDate(today().yearDays, today().year)
}

// days returns the number of days the chain has been unbroken for.
func (c *chain) days() int {
	return c.db.Days()
}

// year returns the year in which the chain was marked.
func (c *chain) year() int {
	return c.db.Year()
}

// broken reports if the chain has been broken.
func (c *chain) broken() bool {
	if c.days() == today().yearDays && c.year() == today().year {
		return false
	}
	if c.days() == yesterday().yearDays && c.year() == yesterday().year {
		return false
	}
	return true
}
