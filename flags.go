package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	dbtc = kingpin.New("dont_break_the_chain", "Keep track of your chain progress from the command line.")

	debug = dbtc.Flag("debug", "Set debug logging level.").
		Bool()

	shortProgress = dbtc.Flag("short-progress", "Display progress as a series of marks rather than on a full calendar.").
			Short('s').
			Bool()
)
