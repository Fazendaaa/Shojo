package shojo

type Package struct {
	Name     string `yaml:"name"`
	Revision string `yaml:"revision"`
}

// remove https://stackoverflow.com/a/37335777/7092954
func remove(slice []Package, s int) (_ []Package) {
	return append(slice[:s], slice[s+1:]...)
}

func isPackagePresent(project Project, packageName string) (present bool) {
	for _, toCheck := range project.Packages {
		if toCheck.Name == packageName {
			return true
		}
	}

	return false
}

func removePackage(project *Project, packageName string) (fail error) {
	toRemove := -1

	for index, toCheck := range project.Packages {
		if toCheck.Name == packageName {
			toRemove = index
			break
		}
	}

	project.Packages = remove(project.Packages, toRemove)

	return fail
}
