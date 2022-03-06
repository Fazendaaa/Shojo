package controllers

func UpgradePackages(projectPath string) (resultChannel chan PackageResponse, fail error) {
	resultChannel = make(chan PackageResponse)

	return resultChannel, fail
}
