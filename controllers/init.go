package controllers

import (
	"fmt"
	"os"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func InitProject() {
	path, fail := os.Getwd()

	if nil != fail {
		fmt.Errorf("%w;\ncould not read current directory", fail)
	}

	shojo.CreateProject(path)
}
