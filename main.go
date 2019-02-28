package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	initLogger()

	// ensure our config file exists or ask the user to create one
	cfg, err := configLocation()
	if err != nil {
		log.Fatal().Err(err)
	}
	if _, err := os.Stat(cfg); os.IsNotExist(err) {
		if newConfigOK(cfg) {
			f, err := os.Create(cfg)
			_ = f.Close()
			if err != nil {
				log.Fatal().Msgf("unable to create config file at %s: %s", cfg, err)
			}
		} else {
			os.Exit(1)
		}
	}

	if len(os.Args) > 1 {
		fmt.Println("sorry, positional parameters are not supported at this time.  we would like to use these to specify different chains in the future.")
	}

	db := dbFile{path: cfg}
	if err = db.Load(); err != nil {
		log.Fatal().Msgf("error loading from database file at %s: %s", cfg, err)
	}

	defaultChain := chain{&db}

	// if db has zero vals we need to do an initial population
	if db.yearDays == 0 && db.year == 0 {
		if err := defaultChain.markToday(); err != nil {
			log.Fatal().Err(err)
		}
	}

	if defaultChain.broken() {
		fmt.Printf("oh no, you broke the chain after %d days. keep this one going longer.\n", defaultChain.days())
	}

}
