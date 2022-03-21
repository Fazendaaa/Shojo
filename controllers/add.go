package controllers

import (
	"fmt"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func AddPackages(packages []string, projectPath string) (resultChannel chan PackageResponse, fail error) {
	project, fail := shojo.ReadProject(projectPath)

	if nil != fail {
		return resultChannel, fmt.Errorf("%w;\nerror while reading project file", fail)
	}

	resultChannel = make(chan PackageResponse)

	for _, packageName := range packages {
		go func(name string) {
			resultChannel <- PackageResponse{
				PackageName: name,
				Response:    project.AddToProject(name),
			}
		}(packageName)
	}

	return resultChannel, fail
}
