package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func createFileFromTemplate(userFullName, projName, templ, dest string, info Project) {
	filename := filepath.Join(SRCPATH, userFullName, projName, dest)
	tempfile := filepath.Join(SRCPATH, userFullName, "gobi", templ)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t, _ := template.ParseFiles(tempfile)
		f, _ := os.Create(filename)
		t.Execute(f, info)
		fileCreated(filepath.Join(userFullName, projName, dest))
	}
}

func cl(user UserConfig, projName string) {
	firstName, secondName := validateProjName(projName)
	userFullName := user.FullName()
	buildDir := filepath.Join(SRCPATH, userFullName, projName)
	checkProjExists(buildDir)
	os.MkdirAll(buildDir, 0744)

	info := Project{projName, userFullName, "cl"}

	createFileFromTemplate(userFullName, projName, "templates/cl/clproj.go.templ", secondName+".go", info)
	createFileFromTemplate(userFullName, firstName, "templates/cl/README.md", "README.md", info)
	createFileFromTemplate(userFullName, firstName, "templates/LICENSE", "LICENSE", info)
	createFileFromTemplate(userFullName, firstName, "templates/VERSION", "VERSION", info)
	creationReady()
}
