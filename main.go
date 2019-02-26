package main

import (
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

	//
}
