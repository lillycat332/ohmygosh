package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
	"github.com/jwalton/gchalk"
)

type Config struct {
	Prompt   string
	Greeting string
}

var configFile string
var configPath string
var homePath string

func init() {
	// read in configuration and set up path variables
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("couldn't open your home directory: %s", err)
	}
	configPath = fmt.Sprintf("%s/lib", homePath)
	configFile = fmt.Sprintf("%s/goshrc.toml", configPath)
	err = os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create config directory: %s", err)
	}
}

func main() {
	var conf Config
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		conf.Greeting = ""
		conf.Prompt = "$"
	} else {
		if _, err := toml.DecodeFile(configFile, &conf); err != nil {
			fmt.Printf("There is an issue with your config file! %s", err)
		}
	}

	fmt.Printf("%s", conf.Greeting)
	// set the user and hostnames
	usr, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := usr.Username
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s %s ", username+"@"+hostname, gchalk.Bold(conf.Prompt))
		// read keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// execute input
		if err = runCommand(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
