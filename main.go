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

	// set database as file and load into db file
	db := dbFile{path: cfg}
	if err = db.Load(); err != nil {
		log.Fatal().Msgf("error loading from database file at %s: %s", cfg, err)
	}

	// make a new chain
	defaultChain := chain{&db}

	// if db has zero vals we need to do an initial population
	if db.LastDate() == "" {
		fmt.Println("oh, I guess this is the first day for this chain. today will be your first of many. kick ass.")
		if err := defaultChain.Start(); err != nil {
			log.Fatal().Msgf("unable to start chain: %s", err)
		}
		if err := defaultChain.MarkToday(); err != nil {
			log.Fatal().Msgf("unable to mark today on the chain: %s", err)
		}
	}

	if defaultChain.Broken() {
		chainLen := defaultChain.Length()
		if chainLen == 1 {
			fmt.Println("oh no, you broke the chain after 1 day. keep this one going longer.")
		} else {
			fmt.Printf("oh no, you broke the chain after %d days. keep this one going longer.\n", chainLen)
		}
		if err := defaultChain.Start(); err != nil {
			log.Fatal().Msgf("unable to start chain: %s", err)
		}
	}

	// mark today as done
	if err := defaultChain.MarkToday(); err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println("today has been marked. great work!")

	// get the new length after marking today
	chainLen := defaultChain.Length()
	// show some motivational messages based on how long the chain has been running
	if chainLen == 1 {
		fmt.Printf("this is your first day. ")
	} else {
		fmt.Printf("you've been at it for %d days now. ", chainLen)
	}
	switch {
	default:
		fmt.Println()
	case chainLen <= 1:
		fmt.Println("every journey has a beginning, set yourself up for success.")
	case chainLen <= 2:
		fmt.Println("come on, you got this.")
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

	// show some X's to give a feel for how long the chain has been going
	fmt.Printf("progress -> ")
	for x := 0; x < chainLen; x++ {
		fmt.Printf("X")
	}
	fmt.Println("")

}
