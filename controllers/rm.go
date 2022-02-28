package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func RmPackages(packages []string, path string) {
	for _, packageName := range packages {
		result, pkgFail := shojo.RmFromProject(path, packageName)

		if nil != pkgFail {
			fmt.Printf(`%v;
error while installing package '%s';
halted execution
`, pkgFail, packageName)
			break
		}

		fmt.Println(result)
	}
}
