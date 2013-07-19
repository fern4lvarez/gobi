package main

import (
	"encoding/json"
	"fmt"
	c "github.com/wsxiaoys/terminal/color"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var exists = false

func TestGobiHelp(t *testing.T) {
	setup()
	defer teardown()
	assertCommand(t, true, "gobi help")
}

func TestGobiVersion(t *testing.T) {
	setup()
	defer teardown()
	assertCommand(t, true, "gobi version")
}

func TestGobiWhoami(t *testing.T) {
	setup()
	defer teardown()
	assertCommand(t, true, "gobi whoami")
}

func TestGobiCl(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	assertCommand(t, true, "gobi cl clapp")
	assertCommand(t, true, "gobi cl clapp/app")
	assertCommand(t, false, "gobi cl clapp")
	assertCommand(t, true, "gobi cl clapp2/app")
}

func TestGobiPkg(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	assertCommand(t, true, "gobi pkg gopkg")
	assertCommand(t, true, "gobi pkg gopkg/pkg")
	assertCommand(t, false, "gobi pkg gopkg")
	assertCommand(t, true, "gobi pkg gopkg2/pkg")
}

func TestGobiWeb(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	assertCommand(t, true, "gobi web goweb")
	assertCommand(t, true, "gobi web goweb/web")
	assertCommand(t, false, "gobi web goweb")
	assertCommand(t, true, "gobi web goweb2/web")
}

func TestGobiMix(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	assertCommand(t, true, "gobi pkg gomix")
	assertCommand(t, true, "gobi web gomix/site")
	assertCommand(t, true, "gobi cl gomix/cli")
	assertCommand(t, true, "gobi pkg gomix/mix")
}

func assertCommand(t *testing.T, b bool, cmd string) {
	c.Println("@{!b} $", cmd)
	cmdSl := strings.Split(cmd, " ")
	out, err := exec.Command(cmdSl[0], cmdSl[1:]...).Output()
	if b {
		if err != nil {
			t.Error()
		}
	} else {
		if err == nil {
			t.Error()
		}
	}
	fmt.Println(string(out))
}

func setup() {
	if _, err := os.Stat(GOBI_CONFIG); err == nil {
		exists = true
		os.Rename(GOBI_CONFIG, GOBI_CONFIG+".tmp")
		createTestConfig()
	} else {
		createTestConfig()
	}
}

func teardown() {
	if exists {
		os.Rename(GOBI_CONFIG+".tmp", GOBI_CONFIG)
	} else {
		os.Remove(GOBI_CONFIG)
	}
	exists = false
}

func createTestConfig() {
	conf := &UserConfig{"Test Name", "testUserName", "github.com", "test@mail.com", "MIT"}
	b, _ := json.Marshal(conf)
	ioutil.WriteFile(GOBI_CONFIG, []byte(b), 0744)
}

func cleanupFiles(path string) {
	os.RemoveAll(path)
}
