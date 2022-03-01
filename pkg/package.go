package shojo

type Package struct {
	Name     string `yaml:"name"`
	Revision string `yaml:"revision"`
}

func isPackagePresent(project Project, packageName string) (present bool) {
	for _, toCheck := range project.Packages {
		if toCheck.Name == packageName {
			return true
		}
	}

	return false
}
