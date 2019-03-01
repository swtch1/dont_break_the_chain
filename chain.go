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

// yearDays returns the yearDays the chain was last marked with.
func (c *chain) yearDays() int {
	return c.db.Days()
}

// year returns the year in which the chain was last marked with.
func (c *chain) year() int {
	return c.db.Year()
}

// Broken reports if the chain has been broken.
func (c *chain) Broken() bool {
	if c.yearDays() == today().yearDays && c.year() == today().year {
		return false
	}
	if c.yearDays() == yesterday().yearDays && c.year() == yesterday().year {
		return false
	}
	return true
}

// Length returns the number of days the chain has been unbroken for.
func (c *chain) Length() int {
	// so here we could do some calculus but it may make more sense to refactor the way we store the date so that we
	// can just bring it in and calculate.  I think we've leaked our logic into the config unnecessarily.
}
