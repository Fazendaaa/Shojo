package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func AddPackage(packageName string, path string) (fail error) {
	_, pkgFail := shojo.AddToProject(path, packageName)

	if nil != pkgFail {
		return fmt.Errorf(`%v;
error while installing package '%s';
halted execution`, pkgFail, packageName)
	}

	return fail
}
