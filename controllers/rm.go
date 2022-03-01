package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func rmFromProject(packageName string, path string) (fail error) {
	mutex.Lock()
	defer mutex.Unlock()

	return shojo.RmFromDescription(path, packageName)
}

func mPackage(packageName string, path string, uninstall bool) (fail error) {
	result, fail := shojo.RmFromProject(path, packageName, uninstall)

	if nil != fail {
		return fmt.Errorf(`%v;
error while removing package '%s';
halted execution due to:
%s`, fail, packageName, result)
	}

	return rmFromProject(packageName, path)
}

func RmPackages(packages []string, projectPath string, uninstall bool) chan PackageResponse {
	resultChannel := make(chan PackageResponse)

	for _, packageName := range packages {
		go func(name string) {
			resultChannel <- PackageResponse{
				PackageName: name,
				Response:    mPackage(name, projectPath, uninstall),
			}
		}(packageName)
	}

	return resultChannel
}
