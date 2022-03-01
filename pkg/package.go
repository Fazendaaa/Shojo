package shojo

type Package struct {
	Name     string `yaml:"name"`
	Revision string `yaml:"revision"`
}

func isPackagePresent(project Project, packageName string) (present bool) {
	return present
}
