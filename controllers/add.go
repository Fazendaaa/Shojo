package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func addToProject(packageName string, path string) (fail error) {
	mutex.Lock()
	defer mutex.Unlock()

	return shojo.AddToDescription(path, packageName)
}

func addPackage(packageName string, path string) (fail error) {
	result, fail := shojo.AddToProject(path, packageName)

	if nil != fail {
		return fmt.Errorf(`%v;
error while installing package '%s';
halted execution due to:
%s`, fail, packageName, result)
	}

	return addToProject(packageName, path)
}

func AddPackages(packages []string, projectPath string) chan PackageResponse {
	resultChannel := make(chan PackageResponse)

	for _, packageName := range packages {
		go func(name string) {
			resultChannel <- PackageResponse{
				PackageName: name,
				Response:    addPackage(name, projectPath),
			}
		}(packageName)
	}

	return resultChannel
}
