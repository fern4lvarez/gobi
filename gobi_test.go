package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	c "github.com/wsxiaoys/terminal/color"
)

var exists = false

func TestGobiWrongCommands(t *testing.T) {
	setupGithub()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "test"))
	assertCommand(t, false, "gobi foo")
	assertCommand(t, false, "gobi bersion")
	assertCommand(t, false, "gobi cl web app")
	assertCommand(t, false, "gobi pkg multiple vars")
	assertCommand(t, false, "gobi web many vars not allowed")
}

func TestGobiHelp(t *testing.T) {
	setupGithub()
	defer teardown()
	assertCommand(t, true, "gobi help")
}

func TestGobiVersion(t *testing.T) {
	setupGithub()
	defer teardown()
	assertCommand(t, true, "gobi version")
}

func TestGobiWhoami(t *testing.T) {
	setupGithub()
	defer teardown()
	assertCommand(t, true, "gobi whoami")
}

func TestGobiCl(t *testing.T) {
	setupGithub()
	assertCommand(t, true, "gobi cl clapp")
	assertCommand(t, true, "gobi cl clapp/app")
	assertCommand(t, false, "gobi cl clapp")
	assertCommand(t, true, "gobi cl clapp2/app")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GITHUB, "test"))

	setupGoogle()
	assertCommand(t, true, "gobi cl clapp")
	assertCommand(t, true, "gobi cl clapp/app")
	assertCommand(t, false, "gobi cl clapp")
	assertCommand(t, true, "gobi cl clapp2/app")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GOOGLE, "p", "clapp"))
	cleanupFiles(filepath.Join(SRCPATH, GOOGLE, "p", "clapp2"))
}

func TestGobiPkg(t *testing.T) {
	setupGithub()
	assertCommand(t, true, "gobi pkg gopkg")
	assertCommand(t, true, "gobi pkg gopkg/pkg")
	assertCommand(t, false, "gobi pkg gopkg")
	assertCommand(t, true, "gobi pkg gopkg2/pkg")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GITHUB, "test"))

	setupGoogle()
	assertCommand(t, true, "gobi pkg gopkg")
	assertCommand(t, true, "gobi pkg gopkg/pkg")
	assertCommand(t, false, "gobi pkg gopkg")
	assertCommand(t, true, "gobi pkg gopkg2/pkg")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GOOGLE, "p", "gopkg"))
	cleanupFiles(filepath.Join(SRCPATH, GOOGLE, "p", "gopkg2"))
}

func TestGobiWeb(t *testing.T) {
	setupGithub()
	assertCommand(t, true, "gobi web goweb")
	assertCommand(t, true, "gobi web goweb/web")
	assertCommand(t, false, "gobi web goweb")
	assertCommand(t, true, "gobi web goweb2/web")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GITHUB, "test"))

	setupGoogle()
	assertCommand(t, true, "gobi web goweb")
	assertCommand(t, true, "gobi web goweb/web")
	assertCommand(t, false, "gobi web goweb")
	assertCommand(t, true, "gobi web goweb2/web")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GOOGLE, "p", "goweb"))
	cleanupFiles(filepath.Join(SRCPATH, GOOGLE, "p", "goweb2"))
}

func TestGobiMix(t *testing.T) {
	setupGithub()
	assertCommand(t, true, "gobi pkg gomix")
	assertCommand(t, true, "gobi web gomix/site")
	assertCommand(t, true, "gobi cl gomix/cli")
	assertCommand(t, true, "gobi pkg gomix/mix")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GITHUB, "test"))

	setupGoogle()
	assertCommand(t, true, "gobi pkg gomix")
	assertCommand(t, true, "gobi web gomix/site")
	assertCommand(t, true, "gobi cl gomix/cli")
	assertCommand(t, true, "gobi pkg gomix/mix")
	teardown()
	cleanupFiles(filepath.Join(SRCPATH, GOOGLE, "p", "gomix"))
}

func assertCommand(t *testing.T, b bool, cmd string) {
	c.Println("@{!b} $", cmd)
	cmdSl := strings.Split(cmd, " ")
	out, err := exec.Command(cmdSl[0], cmdSl[1:]...).Output()
	if b {
		if err != nil {
			t.Error("Error.")
		}
	} else {
		if err == nil {
			t.Error("Error.")
		}
	}
	fmt.Println(string(out))
}

func setupGithub() {
	setup("Test", "test", GITHUB, "test@mail.com", "MIT")
}

func setupGoogle() {
	setup("Test", "test", GOOGLE, "test@mail.com", "MIT")
}

func setup(name, userName, host, email, license string) {
	if _, err := os.Stat(GOBI_CONFIG); err == nil {
		exists = true
		os.Rename(GOBI_CONFIG, GOBI_CONFIG+".tmp")
		createTestConfig(name, userName, host, email, license)
	} else {
		createTestConfig(name, userName, host, email, license)
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

func createTestConfig(name, userName, host, email, license string) {
	conf := &UserConfig{name, userName, host, email, license}
	b, _ := json.Marshal(conf)
	ioutil.WriteFile(GOBI_CONFIG, []byte(b), 0744)
}

func cleanupFiles(path string) {
	os.RemoveAll(path)
}
