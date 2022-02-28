package shojo

type Package struct {
	Name      string     `yaml:"name"`
	Artefacts []Artifact `yaml:"artefacts"`
}

func isPackagePresent(project Project, packageName string) (present bool) {
	return present
}
