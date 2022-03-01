package controllers

import shojo "github.com/Fazendaaa/Shojo/pkg"

func SetRepository(repository string, projectPath string) (fail error) {
	return shojo.SetRepository(repository, projectPath)
}
