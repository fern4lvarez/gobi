package main

import (
	"encoding/json"
	"fmt"
	c "github.com/wsxiaoys/terminal/color"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	msgFail = "%v method fails. Expects %v, returns %v"
	exists  = false
)

func TestGobiHelp(t *testing.T) {
	setup()
	defer teardown()
	c.Println("@{!b} $ gobi help")
	out, err := exec.Command("gobi", "help").Output()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Println(string(out))
}

func TestGobiVersion(t *testing.T) {
	setup()
	defer teardown()
	c.Println("@{!b} $ gobi version")
	out, err := exec.Command("gobi", "version").Output()
	if err != nil {
		t.Errorf("%v", err)
	}
	fmt.Println(string(out))
}

func TestGobiWhoami(t *testing.T) {
	setup()
	defer teardown()
	c.Println("@{!b} $ gobi whoami")
	out, err := exec.Command("gobi", "whoami").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))
}

func TestGobiCl(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	c.Println("@{!b} $ gobi cl clapp")
	out, err := exec.Command("gobi", "cl", "clapp").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi cl clapp/app")
	out, err = exec.Command("gobi", "cl", "clapp/app").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi cl clapp")
	out, err = exec.Command("gobi", "cl", "clapp").Output()
	if err == nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi cl clapp2/app")
	out, err = exec.Command("gobi", "cl", "clapp2/app").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))
}

func TestGobiPkg(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	c.Println("@{!b} $ gobi pkg gopkg")
	out, err := exec.Command("gobi", "pkg", "gopkg").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi pkg gopkg/pkg")
	out, err = exec.Command("gobi", "pkg", "gopkg/pkg").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi pkg gopkg")
	out, err = exec.Command("gobi", "pkg", "gopkg").Output()
	if err == nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi pkg gopkg2/pkg")
	out, err = exec.Command("gobi", "pkg", "gopkg2/pkg").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))
}

func TestGobiWeb(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	c.Println("@{!b} $ gobi web goweb")
	out, err := exec.Command("gobi", "web", "goweb").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi web goweb/web")
	out, err = exec.Command("gobi", "web", "goweb/web").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi web goweb")
	out, err = exec.Command("gobi", "web", "goweb").Output()
	if err == nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi web goweb2/web")
	out, err = exec.Command("gobi", "web", "goweb2/web").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))
}

func TestGobiMix(t *testing.T) {
	setup()
	defer teardown()
	defer cleanupFiles(filepath.Join(SRCPATH, GITHUB, "testUserName"))
	c.Println("@{!b} $ gobi pkg gomix")
	out, err := exec.Command("gobi", "pkg", "gomix").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi web gomix/site")
	out, err = exec.Command("gobi", "web", "gomix/site").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi cl gomix/cli")
	out, err = exec.Command("gobi", "cl", "gomix/cli").Output()
	if err != nil {
		t.Error()
	}
	fmt.Println(string(out))

	c.Println("@{!b} $ gobi pkg gomix/mix")
	out, err = exec.Command("gobi", "pkg", "gomix/mix").Output()
	if err != nil {
		t.Error()
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
