package shojo

type PackageLock struct {
	Package   Package    `yaml:"package"`
	Artefacts []Artifact `yaml:"artefacts"`
}
