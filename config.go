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
	Name    string `json:"name"`
	Id      string `json:"id"`
	Host    string `json:"host"`
	Email   string `json:"email"`
	License string `json:"license"`
}

func NewConfig() *UserConfig {
	var name, username, host, email, license string
	reader := bufio.NewReader(os.Stdin)

	c.Println("@{!y}No configuration found! @bI'd like to know more about you.")
	// Name
	c.Print("@{!b}Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	for !validateName(name) {
		c.Println("@{!y}Please insert your name.")
		c.Print("@{!b}Name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
	}
	// Username
	c.Print("@{!b}Username: ")
	username, _ = reader.ReadString('\n')
	username = strings.TrimSpace(username)
	for !validateUserName(username) {
		reader = bufio.NewReader(os.Stdin)
		c.Println("@{!y}Wrong username, try again.")
		c.Print("@{!b}Username: ")
		username, _ = reader.ReadString('\n')
		username = strings.TrimSpace(username)
	}
	// Host
	c.Print("@{!b}Host @b(github.com, bitbucket.org or code.google.com)@{!b}: ")
	host, _ = reader.ReadString('\n')
	host = strings.TrimSpace(host)
	for !validateHost(strings.TrimSpace(host)) {
		reader = bufio.NewReader(os.Stdin)
		c.Println("@{!y}Invalid host, try again. @yOptions: github.com, bitbucket.org or code.google.com")
		c.Print("@{!b}Host: ")
		host, _ = reader.ReadString('\n')
		host = strings.TrimSpace(host)
	}
	// Email
	c.Print("@{!b}Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	for !validateEmail(email) {
		reader = bufio.NewReader(os.Stdin)
		c.Println("@{!y}Invalid Email address, try again.")
		c.Print("@{!b}Email: ")
		email, _ = reader.ReadString('\n')
		email = strings.TrimSpace(email)
	}
	// License
	c.Print("@{!b}License: ")
	license, _ = reader.ReadString('\n')
	license = strings.TrimSpace(license)
	for !validateLicense(license) {
		reader = bufio.NewReader(os.Stdin)
		c.Println("@{!y}Invalid license, try again. @yOptions: AGPL, Apache, BSD, BSD3-Clause, Eclipse, GPLv2, GPLv3, LGPLv2.1, LGPLv3, MIT, Mozilla, PublicDomain, no-license.")
		c.Print("@{!b}License: ")
		license, _ = reader.ReadString('\n')
		license = strings.TrimSpace(license)
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

func validateName(name string) bool {
	return name != ""
}

func validateUserName(username string) bool {
	return username != "" && !strings.Contains(username, "/") && !strings.Contains(username, " ")
}

func validateHost(host string) bool {
	hosts := []string{GITHUB, BITBUCKET, GOOGLE}
	for i := range hosts {
		if strings.EqualFold(host, hosts[i]) {
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
