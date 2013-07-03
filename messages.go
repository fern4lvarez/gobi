package main

import (
	c "github.com/wsxiaoys/terminal/color"
	"os"
)

var (
	// Command line errors
	wrongNumberOfArguments = "@{!r}Wrong number of arguments, try again."
	wrongArgument          = "@{!r}Wrong argument, try again."
	noProjectName          = "@{!r}You need to specify a name."
	directoryExists        = "@{!y}Oops! Looks like this directory already exists."

	// Help messages
	seeHelp = "@rSee ´gobi help´ for more info."

	// Welcome message
	sayHi = "@bSay hi to @{!r}gobi@b, your new favourite gopher friend!"

	// Help command
	helpCmd = `@bLooks like you need some help:
  @c- @{!y}gobi whoami@w: Tells you who you are, so where are the projects going to be created.
  @c- @{!y}gobi cl <APPNAME>@w: Creates a simple command line app ready to use.
`
)

func welcome() {
	c.Print(sayHi)
	c.Println("@{b!}", logo)
}

func help() {
	welcome()
	c.Println(helpCmd)
}

func commandLineError(msg string) {
	c.Println(msg, seeHelp)
	os.Exit(1)
}

func whoAreYou(user UserConfig) {
	c.Printf("@bYou are @{!g}%s@b.\n", user.FullName())
}

func fileCreated(file string) {
	c.Println("@g Create", file, "...")
}

func creationReady() {
	c.Println("@{!g} Done!")
}
