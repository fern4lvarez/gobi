package main

import (
	"os"
	"path/filepath"
	"text/template"
)

func createFileFromTemplate(userFullName, projName, templ, dest string, info Project) {
	t, _ := template.ParseFiles(templ)
	f, _ := os.Create(filepath.Join(SRCPATH, userFullName, projName, dest))
	t.Execute(f, info)
	fileCreated(filepath.Join(userFullName, projName, dest))
}

func cl(user UserConfig, projName string) {
	userFullName := user.FullName()
	buildDir := filepath.Join(SRCPATH, userFullName, projName)
	info := Project{projName, userFullName, "cl"}

	if err := os.Mkdir(buildDir, 0744); err != nil {
		commandLineError(directoryExists)
	}

	createFileFromTemplate(userFullName, projName, "templates/cl/clproj.go", projName+".go", info)
	createFileFromTemplate(userFullName, projName, "templates/cl/README.md", "README.md", info)
	createFileFromTemplate(userFullName, projName, "templates/LICENSE", "LICENSE", info)
	createFileFromTemplate(userFullName, projName, "templates/VERSION", "VERSION", info)

	creationReady()
}
