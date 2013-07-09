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
			whoAreYou(user)
		case "help":
			help()
		case "cl":
			if l < 3 {
				commandLineError(noProjectName)
			}
			name := os.Args[2]
			proj := NewProject(name, user.FullName(), "cl")
			proj.Create()
		case "pkg":
			if l < 3 {
				commandLineError(noProjectName)
			}
			name := os.Args[2]
			proj := NewProject(name, user.FullName(), "pkg")
			proj.Create()

		default:
			commandLineError(wrongArgument)
		}
	}
}
