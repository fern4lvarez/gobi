package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/kless/datautil/valid"
	c "github.com/wsxiaoys/terminal/color"
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

// setGobiPath where the templates and the version file will be located
func setGobiPath() {
	if os.Getenv("GOBIPATH") != "" {
		GOBIPATH = filepath.Join(SRCPATH, os.Getenv("GOBIPATH"))
	}
}

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
		promptForm["name"]["welcome"],
		promptForm["name"]["error"],
		promptForm["name"]["welcome2"])
	userName = promptField(validateUserName,
		promptForm["userName"]["welcome"],
		promptForm["userName"]["error"],
		promptForm["userName"]["welcome2"])
	host = promptField(validateHost,
		promptForm["host"]["welcome"],
		promptForm["host"]["error"],
		promptForm["host"]["welcome2"])
	email = promptField(validateEmail,
		promptForm["email"]["welcome"],
		promptForm["email"]["error"],
		promptForm["email"]["welcome2"])
	license = promptField(validateLicense,
		promptForm["license"]["welcome"],
		promptForm["license"]["error"],
		promptForm["license"]["welcome2"])

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
	}
	return user
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
	for _, h := range hosts {
		if strings.EqualFold(host, h) {
			return true
		}
	}
	return false
}

// validateEmail: Must have a correct email format
func validateEmail(email string) bool {
	schema := valid.NewSchema(0)

	_, err := valid.Email(schema, email)
	if err != nil {
		return false
	}
	return true
}

// validateLicense: Must be one of the supported licenses
func validateLicense(license string) bool {
	for _, l := range licenses {
		if strings.EqualFold(license, l) {
			return true
		}
	}
	return false
}
