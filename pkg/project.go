package shojo

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Tex struct {
	version string `yaml:"version"`
}

type TLMGR struct {
	version string `yaml:"version"`
}

type Repository struct {
	url string `yaml:"url"`
}

type Package struct {
	name string `yaml:"name"`
}

type Project struct {
	tex        Tex        `yaml:"tex"`
	tlmgr      TLMGR      `yaml:"tlmgr"`
	repository Repository `yaml:"repository"`
	packages   []Package  `yaml:"packages"`
}

func CreateProject(projectPath string) (fail error) {
	yamlData, fail := yaml.Marshal(&Project{
		tex: Tex{
			version: "",
		},
		tlmgr: TLMGR{
			version: "",
		},
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
