package main

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	configFileName = "dontbreakthechain.conf"
)

// configFilePath is the actual path of the configuration file determined by configLocation().
var configFilePath string

// configLocation returns the location of the config file.
func configLocation() (string, error) {
	if configFilePath != "" {
		return configFilePath, nil
	}

	usr, err := user.Current()
	if err != nil {
		log.Debug().Msgf("unable to get current user: %s", err)
		return "", err
	}

	parent := filepath.Join(usr.HomeDir, ".config")
	loc := filepath.Join(parent, configFileName)
	f, err := os.Stat(parent)
	if err == nil && f.IsDir() {
		return loc, nil
	} else {
		loc = filepath.Join(usr.HomeDir, configFileName)
	}

	configFilePath = loc
	return loc, nil
}

// newConfigOK prompts the user to see if they are OK using a new config file.  The return will be
// whether or nor the user agreed.  This ensures the user knows where the config file is going, and also
// that we don't end up thinking we've broken the chain.
func newConfigOK(path string) (confirmed bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("It looks like you haven't setup a config file.  We're going to make one at %s\n", path)
	fmt.Print("Is this OK? (Y/n) ")
	r, _ := reader.ReadString('\n')
	if strings.TrimRight(strings.ToLower(r), "\n") != "y" {
		fmt.Println("ok, I get it.. forget I said anything.")
		return false
	}
	return true
}
