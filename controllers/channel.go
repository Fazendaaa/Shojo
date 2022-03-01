package controllers

import "sync"

type PackageResponse struct {
	PackageName string
	Response    error
}

var mutex sync.Mutex
