package controllers

import (
	"fmt"
	"os"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func RmPackages(packages []string) {
	path, fail := os.Getwd()

	if nil != fail {
		fmt.Errorf("%w;\ncould not read current directory", fail)
	}

	project, fail := shojo.Load(path)

	fmt.Println(project)

	if nil != fail {
		fmt.Errorf("%w;\nmalformed tex definition", fail)
	}
}
