package controllers

import (
	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func InstallProject(path string) (fail error) {
	return shojo.InstallProject(path)
}
