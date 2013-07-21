package main

import (
	c "github.com/wsxiaoys/terminal/color"
	"os"
	"strings"
)

// gobi logo
const logo = `
              _     _ 
             | |   (_)
   __ _  ___ | |__  _ 
  / _' |/ _ \| '_ \| |
 | (_| | (_) | |_) | |
  \__, |\___/|_.__/|_|
   __/ |              
  |___/  
`
// All supported licenses
type Licenses []string

// print all licenses
func (l Licenses) print() string {
	return strings.Join(l, ", ")
}

// global variables used as print messages
var (
	// Command line errors
	wrongNumberOfArguments = "@{!r}Wrong number of arguments, try again."
	wrongArgument          = "@{!r}Wrong argument, try again."
	noProjectName          = "@{!r}You need to specify a name."
	wrongProjectName       = "@{!r}The project name is not valid."
	projectExists          = "@{!y}Oops! Looks like this project already exists."

	// Help messages
	seeHelp = "@rSee ´gobi help´ for more info."

	// Welcome message
	sayHi = "@bSay hi to @{!r}gobi@b, your new favourite gopher friend!"

	// Help command
	helpCmd = `@bLooks like you need some help:
  @c- @{!y}gobi version@w: Shows current version.
  @c- @{!y}gobi whoami@w: Tells you who you are, so where are the projects going to be created.
  @c- @{!y}gobi cl <APPNAME>@{!c}*@w: Creates a command line app ready to use.
  @c- @{!y}gobi pkg <APPNAME>@{!c}*@w: Creates a Go package with a simple test suite and example.
  @c- @{!y}gobi web <APPNAME>@{!c}*@w: Creates a web application ready to deploy.

  @{!c}* @{!y}<APPNAME> @|can have one or two levels and can't be empty. (Examples: ´regexp´, ´fmt´, ´net/http´, ´crypto/md5´)
`
	// Prompted messages on user configuration form
	promptForm = map[string]map[string]string{
		"name": map[string]string{
			"welcome":  "@{!b}Name: ",
			"error":    "@{!y}Please insert your name.",
			"welcome2": "@{!b}Name: "},
		"userName": map[string]string{
			"welcome":  "@{!b}Username: ",
			"error":    "@{!y}Wrong username, try again.",
			"welcome2": "@{!b}Username: "},
		"host": map[string]string{
			"welcome":  "@{!b}Host @b(github.com, bitbucket.org or code.google.com)@{!b}: ",
			"error":    "@{!y}Invalid host, try again. @yOptions: github.com, bitbucket.org or code.google.com",
			"welcome2": "@{!b}Host: "},
		"email": map[string]string{
			"welcome":  "@{!b}Email: ",
			"error":    "@{!y}Invalid Email address, try again.",
			"welcome2": "@{!b}Email: "},
		"license": map[string]string{
			"welcome":  "@{!b}License: ",
			"error":    c.Sprintf("@{!y}Invalid license, try again. @yOptions: %s", licenses.print()),
			"welcome2": "@{!b}License: "},
	}

	// supported licenses
	licenses = Licenses{"AGPL", "Apache", "BSD", "BSD3-Clause", "Eclipse",
		"GPLv2", "GPLv3", "LGPLv2.1", "LGPLv3", "MIT",
		"Mozilla", "PublicDomain", "WTFPL", "no-license"}
)

// welcome message and the logo
func welcome() {
	c.Print(sayHi)
	c.Println("@{!b}", logo)
}

// help printed
func help() {
	welcome()
	c.Println(helpCmd)
}

//commandLineError prints an specific error and exits the program
func commandLineError(msg string) {
	c.Println(msg, seeHelp)
	os.Exit(1)
}

// fileCreated successfully
func fileCreated(file string) {
	c.Println("@g Create", file, "...")
}

// fileExists message
func fileExists(file string) {
	c.Println("@y File", file, "already exists. Skipping.")
}

// assetsCreated successfully
func assetsCreated(file string) {
	c.Println("@g Create assets on", file, "...")
}

// creadtionReady message
func creationReady() {
	c.Println("@{!g} Done!")
}

// showVersion of the program reading the VERSION file
func showVersion() {
	c.Println("@bVersion@{!b}", Version())
}
