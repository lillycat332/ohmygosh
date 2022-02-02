package main

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func runCommand(input string) error {
	errNoHomeFound := errors.New("cd: no directory specified, and the HOME directory is inaccessible")
	input = strings.TrimSuffix(input, "\n")
	// split cmd and args up
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// cd - change directory
		// assign home directory?
		usr, err := user.Current()
		if err != nil {
			return errors.New("fuck")
		}
		homePath := usr.HomeDir

		if len(args) < 2 {
			if err != nil {
				return errNoHomeFound
			} else {
				return os.Chdir(homePath)
			}
		}
		return os.Chdir(args[1])

	case "exit", "quit", "close":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	// set i/o streams
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// execute and return result of command
	return cmd.Run()
}
