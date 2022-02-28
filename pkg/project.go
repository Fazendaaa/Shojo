package shojo

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Project struct {
	filename   string
	Tex        Tex        `yaml:"tex"`
	TLMGR      TLMGR      `yaml:"tlmgr"`
	Repository Repository `yaml:"repository"`
	Packages   []Package  `yaml:"packages"`
}

type MissingAll struct {
	filename string
	Packages []Package `yaml:"packages"`
}

type MissingTex struct {
	filename string
	TLMGR    TLMGR     `yaml:"tlmgr"`
	Packages []Package `yaml:"packages"`
}

type MissingTLMGR struct {
	filename string
	Tex      Tex       `yaml:"tex"`
	Packages []Package `yaml:"packages"`
}

type MissingRepository struct {
	filename string
	Tex      Tex       `yaml:"tex"`
	TLMGR    TLMGR     `yaml:"tlmgr"`
	Packages []Package `yaml:"packages"`
}

func writeProject(data interface{}, filePath string) (fail error) {
	yamlData, fail := yaml.Marshal(data)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while parsing data to be saved into project file", fail)
	}

	fail = ioutil.WriteFile(filePath, yamlData, 0644)

	if fail != nil {
		return fmt.Errorf("%w;\nerror to write data into the file", fail)
	}

	return fail
}

func CreateProject(projectPath string) (fail error) {
	var data interface{}
	filename := "shojo.yaml"
	tlmgrVersion, fail := tlmgrVersion()

	if nil != fail {
		tlmgrVersion = ""
	}

	texVersion, fail := texVersion()

	if nil != fail {
		texVersion = ""
	}

	if "" == tlmgrVersion && "" == texVersion {
		data = &MissingAll{
			filename: filename,
		}
	}
	if "" == texVersion {
		data = &MissingTex{
			filename: filename,
			TLMGR: TLMGR{
				Version: tlmgrVersion,
			},
		}
	}
	if "" == tlmgrVersion {
		data = &MissingTLMGR{
			filename: filename,
			Tex: Tex{
				Version: texVersion,
			},
		}
	}
	if "" != tlmgrVersion && "" != texVersion {
		data = &MissingRepository{
			filename: filename,
			Tex: Tex{
				Version: texVersion,
			},
			TLMGR: TLMGR{
				Version: tlmgrVersion,
			},
		}
	}

	return writeProject(data, filename)
}

func InstallProject(path string) (fail error) {
	return fail
}

func AddToProject(path string, packageName string) (result string, fail error) {
	project, fail := load(path)

	if fail != nil {
		return result, fmt.Errorf("%w;\nerror while reading file from: %s", fail, path)
	}
	if isPackagePresent(project, packageName) {
		return result, fmt.Errorf("package '%s' already present in project", packageName)
	}

	result, fail = installPackage(packageName)

	if fail != nil {
		return result, fmt.Errorf("%w;\nerror while installing '%s' package", fail, path)
	}

	project.Packages = append(project.Packages, Package{
		Name: packageName,
	})
	fail = writeProject(project, project.filename)

	if fail != nil {
		return result, fmt.Errorf("%w;\nerror while updating project file after installing '%s' package", fail, path)
	}

	return result, fail
}

func RmFromProject(path string, packageName string) (result string, fail error) {
	project, fail := load(path)

	if fail != nil {
		return result, fmt.Errorf("%w;\nerror while reading file from: %s", fail, path)
	}
	if !isPackagePresent(project, packageName) {
		return result, fmt.Errorf("package '%s' not present in project", packageName)
	}

	result, fail = removePackage(packageName)

	if fail != nil {
		return result, fmt.Errorf("%w;\nerror while uninstalling '%s' package", fail, path)
	}

	fail = writeProject(project, project.filename)

	if fail != nil {
		return result, fmt.Errorf("%w;\nerror while updating project file after removing '%s' package", fail, path)
	}

	return result, fail
}
