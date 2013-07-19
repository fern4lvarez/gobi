package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kless/validate"
	c "github.com/wsxiaoys/terminal/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Global variables used in the whole application
var (
	GOPATH      = os.Getenv("GOPATH")
	SRCPATH     = filepath.Join(GOPATH, "src")
	HOME        = os.Getenv("HOME")
	GOBI_CONFIG = filepath.Join(HOME, ".gobi.json")
	GITHUB      = "github.com"
	BITBUCKET   = "bitbucket.org"
	GOOGLE      = "code.google.com"
	GOBIPATH    = filepath.Join(SRCPATH, GITHUB, "fern4lvarez", "gobi")
)

// UserConfig contains all information about the current user
type UserConfig struct {
	Name    string `json:"name"`
	Id      string `json:"id"`
	Host    string `json:"host"`
	Email   string `json:"email"`
	License string `json:"license"`
}

// NewConfig promps a form and returns a UserConfig object
// based on the answers
// A JSON file is stored at $HOME/.gobi.json with this content
func NewConfig() *UserConfig {
	c.Println("@{!y}No configuration found! @bI'd like to know more about you.")

	// Prompted user configuration form
	var name, userName, host, email, license string
	name = promptField(validateName,
		"@{!b}Name: ",
		"@{!y}Please insert your name.",
		"@{!b}Name: ")
	userName = promptField(validateUserName,
		"@{!b}Username: ",
		"@{!y}Wrong username, try again.",
		"@{!b}Username: ")
	host = promptField(validateHost,
		"@{!b}Host @b(github.com, bitbucket.org or code.google.com)@{!b}: ",
		"@{!y}Invalid host, try again. @yOptions: github.com, bitbucket.org or code.google.com",
		"@{!b}Host: ")
	email = promptField(validateEmail,
		"@{!b}Email: ",
		"@{!y}Invalid Email address, try again.",
		"@{!b}Email: ")
	license = promptField(validateLicense,
		"@{!b}License: ",
		"@{!y}Invalid license, try again. @yOptions: AGPL, Apache, BSD, BSD3-Clause, Eclipse, GPLv2, GPLv3, LGPLv2.1, LGPLv3, MIT, Mozilla, PublicDomain, no-license.",
		"@{!b}License: ")

	// User config creation
	conf := &UserConfig{name, userName, host, email, license}
	b, _ := json.Marshal(conf)
	ioutil.WriteFile(GOBI_CONFIG, []byte(b), 0744)
	return conf
}

// WhoAreYou pretty prints the UserConfig
func (uc UserConfig) WhoAreYou() {
	c.Printf("@bYou are @{!g}%s @b(@{!g}%s@b)@b. Creating projects on @{!g}%s/%s @bunder @{!g}%s @blicense.\n",
		uc.Name, uc.Email, uc.Host, uc.Id, uc.License)
}

// checkConfig if JSON config file exists and returns the UserConfig if so
func checkConfig() UserConfig {
	var user UserConfig
	b, errRead := ioutil.ReadFile(GOBI_CONFIG)
	if errUnmarshal := json.Unmarshal(b, &user); errRead != nil || errUnmarshal != nil {
		return *NewConfig()
	} else {
		return user
	}
	return UserConfig{}
}

// promptField to validate and save input value
func promptField(validateFunc func(string) bool, welcomeMsg, errorMsg, welcome2Msg string) (resp string) {
	reader := bufio.NewReader(os.Stdin)
	c.Print(welcomeMsg)
	resp, _ = reader.ReadString('\n')
	resp = strings.TrimSpace(resp)
	for !validateFunc(resp) {
		c.Println(errorMsg)
		c.Print(welcome2Msg)
		resp, _ = reader.ReadString('\n')
		resp = strings.TrimSpace(resp)
	}
	return
}

// validateName: Cannot be empty
func validateName(name string) bool {
	return name != ""
}

// validateUserName: Cannot be empty, only one path level
func validateUserName(username string) bool {
	return username != "" && !strings.Contains(username, "/") && !strings.Contains(username, " ")
}

// validateHost: Can only be GITHUB, BITBUCKET OR GOOGLE
func validateHost(host string) bool {
	hosts := []string{GITHUB, BITBUCKET, GOOGLE}
	for i := range hosts {
		if strings.EqualFold(host, hosts[i]) {
			return true
		}
	}
	return false
}

// validateEmail: Must have a correct email format
func validateEmail(email string) bool {
	exp, err := regexp.Compile(validate.RE_BASIC_EMAIL)
	if err != nil {
		fmt.Println(err)
	}
	return exp.MatchString(email)
}

// validateLicense: Must be one of the supported licenses
func validateLicense(license string) bool {
	licenses := []string{"AGPL", "Apache", "BSD", "BSD3-Clause", "Eclipse",
		"GPLv2", "GPLv3", "LGPLv2.1", "LGPLv3",
		"MIT", "Mozilla", "PublicDomain", "no-license"}
	for i := range licenses {
		if strings.EqualFold(license, licenses[i]) {
			return true
		}
	}
	return false
}
