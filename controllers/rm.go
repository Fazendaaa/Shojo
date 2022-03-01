package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func RmPackages(packageName string, path string, uninstall bool) (fail error) {
	_, pkgFail := shojo.RmFromProject(path, packageName, uninstall)

	if nil != pkgFail {
		return fmt.Errorf(`%v;
error while removing package '%s';
halted execution`, pkgFail, packageName)
	}

	return fail
}
