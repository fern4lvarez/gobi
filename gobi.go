/*
gobi is a command line tool that will make your Go development just faster.
It might stand for *Go Boilerplater Injector*, but just think of it as a fun, tiny tool to create and manage quickly your applications written in Go.
*/
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
		setGobiPath()
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
