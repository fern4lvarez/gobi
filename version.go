package main

import (
	"io/ioutil"
	"path/filepath"
)

// Version of the package
func Version() string {
	b, _ := ioutil.ReadFile(filepath.Join(GOBIPATH, "VERSION"))
	v := string(b)
	return v
}
