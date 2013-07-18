package main

import (
	cp "github.com/opesun/copyrecur"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

type Project struct {
	Name       string
	FirstName  string
	SecondName string
	GoGetName  string
	UserId     string
	UserName   string
	UserEmail  string
	Host       string
	License    string
	Typ        string
}

func NewProject(name, typ string, user UserConfig) *Project {
	firstName, secondName := ValidateName(name)
	goGetName := GoGetName(user.Host, user.Id, name)
	return &Project{name, firstName, secondName, goGetName, user.Id, user.Name, user.Email, user.Host, user.License, typ}
}

func (proj Project) Create() {
	if proj.Exists() {
		commandLineError(projectExists)
	}
	switch typ := proj.Typ; typ {
	case "cl":
		proj.Cl()
	case "pkg":
		proj.Pkg()
	case "web":
		proj.Web()
	}
	creationReady()
}

func (proj Project) Cl() {
	var buildDir, buildDirFirst string
	if proj.Host == GOOGLE {
		buildDir = filepath.Join(SRCPATH, proj.Host, "p", proj.Name)
		buildDirFirst = filepath.Join(SRCPATH, proj.Host, "p", proj.FirstName)
	} else {
		buildDir = filepath.Join(SRCPATH, proj.Host, proj.UserId, proj.Name)
		buildDirFirst = filepath.Join(SRCPATH, proj.Host, proj.UserId, proj.FirstName)
	}
	os.MkdirAll(buildDir, 0744)
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "AUTHORS"), "AUTHORS.tpl")
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "VERSION"), "VERSION.tpl")
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "LICENSE"),
		filepath.Join("license", proj.License+".tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "README.md"),
		filepath.Join(proj.Typ, "README.md.tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDir, proj.SecondName+".go"),
		filepath.Join(proj.Typ, "proj.go.tpl"))
}

func (proj Project) Pkg() {
	var buildDir, buildDirFirst string
	if proj.Host == GOOGLE {
		buildDir = filepath.Join(SRCPATH, proj.Host, "p", proj.Name)
		buildDirFirst = filepath.Join(SRCPATH, proj.Host, "p", proj.FirstName)
	} else {
		buildDir = filepath.Join(SRCPATH, proj.Host, proj.UserId, proj.Name)
		buildDirFirst = filepath.Join(SRCPATH, proj.Host, proj.UserId, proj.FirstName)
	}
	os.MkdirAll(buildDir, 0744)
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "AUTHORS"), "AUTHORS.tpl")
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "VERSION"), "VERSION.tpl")
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "LICENSE"),
		filepath.Join("license", proj.License+".tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "README.md"),
		filepath.Join(proj.Typ, "README.md.tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDir, proj.SecondName+".go"),
		filepath.Join(proj.Typ, "proj.go.tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDir, proj.SecondName+"_test.go"),
		filepath.Join(proj.Typ, "proj_test.go.tpl"))
	os.MkdirAll(filepath.Join(buildDirFirst, "examples"), 0744)
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "examples", proj.SecondName+"_example.go"),
		filepath.Join(proj.Typ, "example.go.tpl"))
}

func (proj Project) Web() {
	var buildDir, buildDirFirst, staticDir string
	if proj.Host == GOOGLE {
		buildDir = filepath.Join(SRCPATH, proj.Host, "p", proj.Name)
		buildDirFirst = filepath.Join(SRCPATH, proj.Host, "p", proj.FirstName)
	} else {
		buildDir = filepath.Join(SRCPATH, proj.Host, proj.UserId, proj.Name)
		buildDirFirst = filepath.Join(SRCPATH, proj.Host, proj.UserId, proj.FirstName)
	}
	staticDir = filepath.Join(buildDir, "static")
	os.MkdirAll(staticDir, 0744)
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "AUTHORS"), "AUTHORS.tpl")
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "VERSION"), "VERSION.tpl")
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "LICENSE"),
		filepath.Join("license", proj.License+".tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "README.md"),
		filepath.Join(proj.Typ, "README.md.tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDir, proj.SecondName+".go"),
		filepath.Join(proj.Typ, "proj.go.tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, ".godir"),
		filepath.Join(proj.Typ, "godir.tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDirFirst, "Procfile"),
		filepath.Join(proj.Typ, "Procfile.tpl"))
	proj.CreateFileFromTemplate(filepath.Join(buildDir, "index.html"),
		filepath.Join(proj.Typ, "index.html.tpl"))
	CopyAssets(filepath.Join(GOBIPATH, "templates", "web"), staticDir)
}

func (proj Project) Exists() bool {
	var err error
	if proj.Host == GOOGLE {
		_, err = os.Stat(filepath.Join(SRCPATH, proj.Host, "p", proj.Name))
	} else {
		_, err = os.Stat(filepath.Join(SRCPATH, proj.Host, proj.UserId, proj.Name))
	}
	return err == nil
}

func (proj Project) CreateFileFromTemplate(file, temp string) {
	tempfile := filepath.Join(GOBIPATH, "templates", temp)
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t, _ := template.ParseFiles(tempfile)
		f, _ := os.Create(file)
		t.Execute(f, proj)
		fileCreated(file)
	} else {
		fileExists(file)
	}
}

func GoGetName(host, userid, name string) string {
	if host == GOOGLE {
		return filepath.Join(host, "p", name)
	} else {
		return filepath.Join(host, userid, name)
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

func CopyAssets(src, dest string) {
	cp.CopyDir(filepath.Join(src, "css"),
		filepath.Join(dest, "css"))
	cp.CopyDir(filepath.Join(src, "js"),
		filepath.Join(dest, "js"))
	cp.CopyDir(filepath.Join(src, "img"),
		filepath.Join(dest, "img"))
	assetsCreated(dest)
}
