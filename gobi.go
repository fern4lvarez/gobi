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
		case "help":
			help()
		case "cl":
			if l < 3 {
				commandLineError(noProjectName)
			}
			name := os.Args[2]
			proj := NewProject(name, "cl", user)
			proj.Create()
		case "pkg":
			if l < 3 {
				commandLineError(noProjectName)
			}
			name := os.Args[2]
			proj := NewProject(name, "pkg", user)
			proj.Create()
		case "web":
			if l < 3 {
				commandLineError(noProjectName)
			}
			name := os.Args[2]
			proj := NewProject(name, "web", user)
			proj.Create()

		default:
			commandLineError(wrongArgument)
		}
	}
}
