package shojo

import (
	"bytes"
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

func (project Project) getFilename() string {
	return project.filename
}

type MissingAll struct {
	filename string
	Packages []Package `yaml:"packages"`
}

func (project MissingAll) getFilename() string {
	return project.filename
}

type MissingTex struct {
	filename string
	TLMGR    TLMGR     `yaml:"tlmgr"`
	Packages []Package `yaml:"packages"`
}

func (project MissingTex) getFilename() string {
	return project.filename
}

type MissingTLMGR struct {
	filename string
	Tex      Tex       `yaml:"tex"`
	Packages []Package `yaml:"packages"`
}

func (project MissingTLMGR) getFilename() string {
	return project.filename
}

type MissingRepository struct {
	filename string
	Tex      Tex       `yaml:"tex"`
	TLMGR    TLMGR     `yaml:"tlmgr"`
	Packages []Package `yaml:"packages"`
}

func (project MissingRepository) getFilename() string {
	return project.filename
}

type GenericRepository interface {
	getFilename() string
}

func writeToProject(data GenericRepository) (fail error) {
	yamlData, fail := yaml.Marshal(&data)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while parsing data to be saved into project file", fail)
	}

	buffer := bytes.NewBuffer(yamlData)
	yamlEncoder := yaml.NewEncoder(buffer)

	yamlEncoder.SetIndent(2)

	fail = ioutil.WriteFile(data.getFilename(), buffer.Bytes(), 0644)

	if fail != nil {
		return fmt.Errorf("%w;\nerror to write data into the file", fail)
	}

	return fail
}

func CreateProject(projectPath string) (fail error) {
	var data GenericRepository
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

	return writeToProject(data)
}

func InstallProject(path string) (fail error) {
	project, fail := load(path)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while reading file from: %s", fail, path)
	}

	if 0 == len(project.Packages) {
		return fmt.Errorf("no packages to install listed in shojo's config in: %s", path)
	}

	for _, toInstall := range project.Packages {
		installed, fail := isInstalled(toInstall.Name)

		if fail != nil {
			return fmt.Errorf(`%w;
error while checking whether or not '%s' package is installed;`, fail, toInstall.Name)
		}

		if !installed {
			result, fail := installPackage(toInstall.Name)

			if fail != nil {
				return fmt.Errorf(`%s;
%w;
error while installing '%s' package;`, result, fail, toInstall.Name)
			}
		}
	}

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

	return result, fail
}

func AddToDescription(path string, packageName string) (fail error) {
	project, fail := load(path)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while reading file from: %s", fail, path)
	}

	show, fail := packageShow(packageName)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while fetching information about installed '%s' package", fail, path)
	}

	project.Packages = append(project.Packages, Package{
		Name:     packageName,
		Revision: show.Revision,
	})

	return writeToProject(project)
}

func RmFromProject(path string, packageName string, uninstall bool) (result string, fail error) {
	project, fail := load(path)

	if fail != nil {
		return result, fmt.Errorf("%w;\nerror while reading file from: %s", fail, path)
	}
	if !isPackagePresent(project, packageName) {
		return result, fmt.Errorf("package '%s' not present in project", packageName)
	}

	if uninstall {
		result, fail = uninstallPackage(packageName)

		if fail != nil {
			return result, fmt.Errorf("%w;\nerror while uninstalling '%s' package", fail, path)
		}
	}

	return result, fail
}

func RmFromDescription(path string, packageName string) (fail error) {
	project, fail := load(path)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while reading file from: %s", fail, path)
	}

	fail = removePackage(&project, packageName)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while removing '%s' package", fail, path)
	}

	fail = writeToProject(project)

	if fail != nil {
		return fmt.Errorf("%w;\nerror while updating project file after removing '%s' package", fail, path)
	}

	return fail
}
