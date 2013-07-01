package main

import (
	"encoding/json"
	"fmt"
	c "github.com/wsxiaoys/terminal/color"
	"io/ioutil"
	"os"
	"path/filepath"
)

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

var (
	GOPATH      = os.Getenv("GOPATH")
	SRCPATH     = filepath.Join(GOPATH, "src")
	HOME        = os.Getenv("HOME")
	GOBI_CONFIG = filepath.Join(HOME, ".gobi.json")
)

func createConfig() UserConfig {
	var user string
	var prov string
	c.Println("@{!y}No configuration found! @bI'd like to know more about you.")
	c.Print("@{!b}Username: ")
	fmt.Scanf("%s", &user)
	c.Print("@{!b}Provider (github.com, bitbucket.com, local...): ")
	fmt.Scanf("%s", &prov)
	conf := UserConfig{user, prov}
	b, _ := json.Marshal(conf)
	ioutil.WriteFile(GOBI_CONFIG, []byte(b), 0744)
	return conf
}

func checkConfig() UserConfig {
	var user UserConfig
	b, errRead := ioutil.ReadFile(GOBI_CONFIG)
	if errUnmarshal := json.Unmarshal(b, &user); errRead != nil || errUnmarshal != nil {
		return createConfig()
	} else {
		return user
	}
	return UserConfig{"NoUser", "NoProvider"}
}

type UserConfig struct {
	Name     string
	Provider string
}

func (uc UserConfig) FullName() string {
	return uc.Provider + "/" + uc.Name
}
