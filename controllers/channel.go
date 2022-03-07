package controllers

import (
	shojo "github.com/Fazendaaa/Shojo/pkg"
)

type PackageResponse struct {
	PackageName string
	Response    error
	Project     *shojo.Project
}
