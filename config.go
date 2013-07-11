package main

import (
	"encoding/json"
	"fmt"
	"github.com/kless/validate"
	c "github.com/wsxiaoys/terminal/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

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

type UserConfig struct {
	Name    string
	Id      string
	Host    string
	Email   string
	License string
}

func NewConfig() *UserConfig {
	var name, username, host, email, license string
	c.Println("@{!y}No configuration found! @bI'd like to know more about you.")
	// Name
	c.Print("@{!b}Name @{!b}: ")
	fmt.Scanf("%s", &name)
	for !validateUserName(name) {
		c.Println("@{!y}Please insert your name.")
		c.Print("@{!b}Name: ")
		fmt.Scanf("%s", &name)
	}
	// Username
	c.Print("@{!b}Username: ")
	fmt.Scanf("%s", &username)
	for !validateUserName(username) {
		c.Println("@{!y}Please insert your username.")
		c.Print("@{!b}Username: ")
		fmt.Scanf("%s", &username)
	}
	// Host
	c.Print("@{!b}Host @b(github.com, bitbucket.org or code.google.com)@{!b}: ")
	fmt.Scanf("%s", &host)
	for !validateHost(host) {
		c.Println("@{!y}Invalid host, try again. @yOptions: github.com, bitbucket.org or code.google.com")
		c.Print("@{!b}Host: ")
		fmt.Scanf("%s", &host)
	}
	// Email
	c.Print("@{!b}Email@{!b}: ")
	fmt.Scanf("%s", &email)
	for !validateEmail(email) {
		c.Println("@{!y}Invalid Email address, try again.")
		c.Print("@{!b}Email: ")
		fmt.Scanf("%s", &email)
	}
	// License
	c.Print("@{!b}License@{!b}: ")
	fmt.Scanf("%s", &license)
	for !validateLicense(license) {
		c.Println("@{!y}Invalid license, try again. @yOptions: MIT, BSD, GPL, LGPL, CC")
		c.Print("@{!b}License: ")
		fmt.Scanf("%s", &license)
	}
	// User config creation
	conf := &UserConfig{name, username, host, email, license}
	b, _ := json.Marshal(conf)
	ioutil.WriteFile(GOBI_CONFIG, []byte(b), 0744)
	return conf
}

func (uc UserConfig) FullName() string {
	return uc.Host + "/" + uc.Name
}

func (uc UserConfig) WhoAreYou() {
	c.Printf("@bYou are @{!g}%s @b(@{!g}%s@b)@b. Creating projects on @{!g}%s/%s @bunder @{!g}%s @blicense.\n",
		uc.Name, uc.Email, uc.Host, uc.Id, uc.License)
}

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

func validateUserName(username string) bool {
	return username != ""
}

func validateHost(host string) bool {
	hosts := []string{GITHUB, BITBUCKET, GOOGLE}
	for i := range hosts {
		if host == hosts[i] {
			return true
		}
	}
	return false
}

func validateEmail(email string) bool {
	exp, err := regexp.Compile(validate.RE_BASIC_EMAIL)
	if err != nil {
		fmt.Println(err)
	}
	return exp.MatchString(email)
}

func validateLicense(license string) bool {
	licenses := []string{"MIT", "BSD", "GPL", "LGPL", "CC"}
	for i := range licenses {
		if license == licenses[i] {
			return true
		}
	}
	return false
}
