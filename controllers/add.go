package controllers

import (
	"fmt"
	"os"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func AddPackage(packages []string) {
	path, fail := os.Getwd()

	if nil != fail {
		fmt.Printf("%v;\ncould not read current directory", fail)
	}

	project, fail := shojo.Load(path)

	if nil != fail {
		fmt.Printf("%v;\nmalformed project file definition", fail)
	}

	for _, packageName := range packages {
		result, pkgFail := shojo.InstallPackage(packageName)

		if nil != pkgFail {
			fmt.Printf(`%v;
error while installing package '%s';
halted execution`, pkgFail, packageName)
			fmt.Println(project)
			break
		}

		fmt.Println(result)
	}
}
