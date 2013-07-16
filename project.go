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
	UId        string
	UName      string
	UEmail     string
	Host       string
	License    string
	Typ        string
}

func NewProject(name, typ string, user UserConfig) *Project {
	firstName, secondName := ValidateName(name)
	return &Project{name, firstName, secondName, user.Id, user.Name, user.Email, user.Host, user.License, typ}
}

func (proj Project) Create() {
	buildDir := filepath.Join(SRCPATH, proj.Host, proj.UId, proj.Name)
	if proj.Exists() {
		commandLineError(projectExists)
	}
	os.MkdirAll(buildDir, 0744)
	createFileFromTemplate(proj.Name, "templates/"+proj.Typ+"/proj.go.tpl", proj.SecondName+".go", proj)
	createFileFromTemplate(proj.FirstName, "templates/"+proj.Typ+"/README.md.tpl", "README.md", proj)
	createFileFromTemplate(proj.FirstName, "templates/LICENSE.tpl", "LICENSE", proj)
	createFileFromTemplate(proj.FirstName, "templates/VERSION.tpl", "VERSION", proj)
	createFileFromTemplate(proj.FirstName, "templates/AUTHORS.tpl", "AUTHORS", proj)
	creationReady()
}

func (proj Project) Exists() bool {
	_, err := os.Stat(filepath.Join(SRCPATH, proj.Host, proj.UId, proj.Name))
	return err == nil
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

func createFileFromTemplate(projName, temp, dest string, proj Project) {
	filename := filepath.Join(SRCPATH, proj.Host, proj.UId, projName, dest)
	tempfile := filepath.Join(GOBIPATH, temp)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t, _ := template.ParseFiles(tempfile)
		f, _ := os.Create(filename)
		t.Execute(f, proj)
		fileCreated(filepath.Join(proj.Host, proj.UId, projName, dest))
	}
}
