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
		switch first := os.Args[1]; first {
		case "whoami":
			user.WhoAreYou()
		case "v", "version":
			showVersion()
		case "help":
			help()
		case "cl", "pkg", "web":
			if l < 3 {
				commandLineError(noProjectName)
			}
			name := os.Args[2]
			proj := NewProject(name, first, user)
			proj.Create()
		default:
			commandLineError(wrongArgument)
		}
	}
}
