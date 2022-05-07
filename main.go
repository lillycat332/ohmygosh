package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	// print the motd
	print(motd())

	// set the user and hostnames ig
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
		fmt.Printf("%s $ ", username+"@"+hostname)
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
