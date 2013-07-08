package main

import (
	"os"
	"regexp"
)

type Project struct {
	Name string
	User string
	Type string
}

func checkProjExists(buildDir string) {
	if _, err := os.Stat(buildDir); err == nil {
		commandLineError(projectExists)
	}
}

func validateProjName(projName string) (firstName string, secondName string) {
	partsProjName := parseProjName(projName)
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

func parseProjName(projName string) []string {
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
