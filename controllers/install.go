package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func InstallProject(path string) (fail error) {
	project, fail := shojo.ReadProject(path)

	if nil != fail {
		return fmt.Errorf("%w;\nerror while reading project file while trying to install its dependencies", fail)
	}

	return project.InstallProject()
}
