package shojo

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Tex struct {
	Version string `yaml:"version"`
}

type TLMGR struct {
	Version string `yaml:"version"`
}

type Repository struct {
	URL string `yaml:"url"`
}

type Package struct {
	Name string `yaml:"name"`
}

type Project struct {
	Tex        Tex        `yaml:"tex"`
	TLMGR      TLMGR      `yaml:"tlmgr"`
	Repository Repository `yaml:"repository"`
	Packages   []Package  `yaml:"packages"`
}

func CreateProject(projectPath string) (fail error) {
	tlmgrVersion, fail := tlmgrVersion()

	if nil != fail {
		tlmgrVersion = ""
	}

	texVersion, fail := texVersion()

	if nil != fail {
		texVersion = ""
	}

	yamlData, fail := yaml.Marshal(&Project{
		Tex: Tex{
			Version: texVersion,
		},
		TLMGR: TLMGR{
			Version: tlmgrVersion,
		},
		Repository: DEFAULT_REPOSITORY,
	})

	if nil != fail {
		return fmt.Errorf("%w;\nerror while generating project file", fail)
	}

	fileName := "shojo.yaml"
	fail = ioutil.WriteFile(fileName, yamlData, 0644)

	if fail != nil {
		return fmt.Errorf("%w;\nerror to write data into the file", fail)
	}

	return fail
}
