package main

import (
	"os"
)

func main() {
	var user UserConfig
	if l := len(os.Args); l == 1 {
		welcome()
		checkConfig()
	} else if l > 3 {
		commandLineError(wrongNumberOfArguments)
	} else {
		user = checkConfig()
		if first := os.Args[1]; first == "whoami" {
			whoAreYou(user)
		} else if first == "help" {
			help()
		} else if first == "cl" {
			if l < 3 {
				commandLineError(noProjectName)
			}
			name := os.Args[2]
			cl(user, name)
		} else {
			commandLineError(wrongNumberOfArguments)
		}
	}
}
