package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// poor mans arg parsing
	var level zerolog.Level
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--debug":
			level = zerolog.DebugLevel
		default:
			level = zerolog.ErrorLevel
		}
	} else {
		level = zerolog.ErrorLevel
	}

	initLogger(level)

	// ensure our config file exists or ask the user to create one
	cfg, err := configLocation()
	if err != nil {
		log.Fatal().Err(err)
	}
	if _, err := os.Stat(cfg); os.IsNotExist(err) {
		if newConfigOK(cfg) {
			log.Debug().Msgf("creating new config file at %s", cfg)
			f, err := os.Create(cfg)
			_ = f.Close()
			if err != nil {
				log.Fatal().Msgf("unable to create config file at %s: %s", cfg, err)
			}
		} else {
			os.Exit(1)
		}
	}

	db := dbFile{path: cfg}
	if err = db.Load(); err != nil {
		log.Fatal().Msgf("error loading from database file at %s: %s", cfg, err)
	}

	defaultChain := chain{&db}

	// if db has zero vals we need to do an initial population
	if db.Date() == "" {
		fmt.Println("oh, I guess this is the first day for this chain. today will be your first of many. kick ass.")
		if err := defaultChain.markToday(); err != nil {
			log.Fatal().Msgf("unable to mark today on the chain: %s", err)
		}
	}

	chainLen, err := defaultChain.Length()
	if err != nil {
		log.Fatal().Err(err)
	}
	if defaultChain.Broken() {
		fmt.Printf("oh no, you broke the chain after %d yearDays. keep this one going longer.\n", chainLen)
	}

	if err := defaultChain.markToday(); err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println("today has been marked. great work!")

	chainLen, err = defaultChain.Length()
	if err != nil {
		log.Fatal().Msgf("unable to get chain length: %s", err)
	}
	fmt.Printf("you've been at it for %d days now. ", chainLen)
	switch {
	default:
		fmt.Println()
	case chainLen <= 1:
		fmt.Println("every journey has a beginning, set yourself up for success.")
	case chainLen <= 2:
		fmt.Println("come one, you got this.")
	case chainLen <= 3:
		fmt.Println("the first few days are the most important, let's go!")
	case chainLen <= 5:
		fmt.Println("making some great progress, do it again tomorrow.")
	case chainLen <= 10:
		fmt.Println("you're on the right track, don't stop now.")
	case chainLen <= 20:
		fmt.Println("wow, you're really getting up there. excellent.")
	case chainLen <= 30:
		fmt.Println("it looks like this could one could go for a while. you're a pro.")
	case chainLen <= 40:
		fmt.Println("shoot for the stars, you've almost touched one at this point.")
	case chainLen <= 50:
		fmt.Println("can you believe it? this is chain is unbelievable!!")
	case chainLen > 50:
		fmt.Println("just killing it.. every day, no matter what.")
	}
}
