package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	dbtc = kingpin.New("dont_break_the_chain", "Keep track of your chain progress from the command line.")

	debug = dbtc.Flag("debug", "Set debug logging level.").
		Bool()
)
