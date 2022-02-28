package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func InitProject(projectPath string) {
	fail := shojo.CreateProject(projectPath)

	if nil != fail {
		fmt.Printf("%v;\ncould not create project in the given directory", fail)
	}
}
