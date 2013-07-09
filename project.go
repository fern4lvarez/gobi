package main

import (
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

type Project struct {
	Name string
	User string
	Typ  string
}

func NewProject(name, user, typ string) *Project {
	return &Project{name, user, typ}
}

func (proj Project) Create() {
	firstName, secondName := proj.ValidateName()
	buildDir := filepath.Join(SRCPATH, proj.User, proj.Name)
	if proj.Exists() {
		commandLineError(projectExists)
	}
	os.MkdirAll(buildDir, 0744)
	createFileFromTemplate(proj.User, proj.Name, "templates/"+proj.Typ+"/proj.go.templ", secondName+".go", proj)
	createFileFromTemplate(proj.User, firstName, "templates/"+proj.Typ+"/README.md", "README.md", proj)
	createFileFromTemplate(proj.User, firstName, "templates/LICENSE", "LICENSE", proj)
	createFileFromTemplate(proj.User, firstName, "templates/VERSION", "VERSION", proj)
	creationReady()
}

func (proj Project) Exists() bool {
	if _, err := os.Stat(filepath.Join(SRCPATH, proj.User, proj.Name)); err == nil {
		return true
	} else {
		return false
	}
}

func (proj Project) ValidateName() (firstName string, secondName string) {
	partsProjName := proj.ParseName()
	if l := len(partsProjName); l == 0 || l > 2 {
		commandLineError(wrongProjectName)
	} else if l == 1 {
		firstName = proj.Name
		secondName = proj.Name
	} else {
		if partsProjName[0] == "" || partsProjName[1] == "" {
			commandLineError(wrongProjectName)
		}
		firstName = partsProjName[0]
		secondName = partsProjName[1]
	}
	return
}

func (proj Project) ParseName() []string {
	delimeter := "/"
	if proj.Name == "" {
		return make([]string, 0)
	}
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(proj.Name, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = proj.Name[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = proj.Name[laststart:len(proj.Name)]
	return result
}

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
