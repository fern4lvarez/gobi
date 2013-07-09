package main

import (
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

type Project struct {
	Name       string
	FirstName  string
	SecondName string
	User       string
	Typ        string
}

func NewProject(name, user, typ string) *Project {
	firstName, secondName := ValidateName(name)
	return &Project{name, firstName, secondName, user, typ}
}

func (proj Project) Create() {
	buildDir := filepath.Join(SRCPATH, proj.User, proj.Name)
	if proj.Exists() {
		commandLineError(projectExists)
	}
	os.MkdirAll(buildDir, 0744)
	createFileFromTemplate(proj.User, proj.Name, "templates/"+proj.Typ+"/proj.go.tpl", proj.SecondName+".go", proj)
	createFileFromTemplate(proj.User, proj.FirstName, "templates/"+proj.Typ+"/README.md.tpl", "README.md", proj)
	createFileFromTemplate(proj.User, proj.FirstName, "templates/LICENSE.tpl", "LICENSE", proj)
	createFileFromTemplate(proj.User, proj.FirstName, "templates/VERSION.tpl", "VERSION", proj)
	creationReady()
}

func (proj Project) Exists() bool {
	if _, err := os.Stat(filepath.Join(SRCPATH, proj.User, proj.Name)); err == nil {
		return true
	} else {
		return false
	}
}

func ParseName(projName string) []string {
	delimeter := "/"
	if projName == "" {
		return make([]string, 0)
	}
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(projName, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = projName[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = projName[laststart:len(projName)]
	return result
}

func ValidateName(projName string) (firstName string, secondName string) {
	partsProjName := ParseName(projName)
	if l := len(partsProjName); l == 0 || l > 2 {
		commandLineError(wrongProjectName)
	} else if l == 1 {
		firstName = projName
		secondName = projName
	} else {
		if partsProjName[0] == "" || partsProjName[1] == "" {
			commandLineError(wrongProjectName)
		}
		firstName = partsProjName[0]
		secondName = partsProjName[1]
	}
	return
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
