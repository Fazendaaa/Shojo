package controllers

import (
	"fmt"
	"os"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func AddPackage(packages []string) {
	path, fail := os.Getwd()

	if nil != fail {
		fmt.Errorf("%w;\ncould not read current directory", fail)
	}

	_, fail = shojo.Load(path)

	if nil != fail {
		fmt.Errorf("%w;\nmalformed tex definition", fail)
	}
}
