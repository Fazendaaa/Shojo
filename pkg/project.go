package shojo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/dgraph-io/badger/v3"
	"gopkg.in/yaml.v3"
)

type Project struct {
	database   *badger.DB
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

var mutex sync.Mutex

func writeToProject(data GenericRepository) (fail error) {
	yamlData, fail := yaml.Marshal(&data)

	if nil != fail {
		return fmt.Errorf("%w;\nerror while parsing data to be saved into project file", fail)
	}

	buffer := bytes.NewBuffer(yamlData)
	yamlEncoder := yaml.NewEncoder(buffer)

	yamlEncoder.SetIndent(2)

	mutex.Lock()
	defer mutex.Unlock()

	fail = ioutil.WriteFile(data.getFilename(), buffer.Bytes(), 0644)

	if nil != fail {
		return fmt.Errorf("%w;\nerror to write data into the file", fail)
	}

	return fail
}

func (project *Project) writeProject() (fail error) {
	return writeToProject(project)
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

func (project *Project) InstallProject() (fail error) {
	if 0 == len(project.Packages) {
		return fmt.Errorf("no packages to install listed in shojo's config in: %s", project.filename)
	}

	for _, toInstall := range project.Packages {
		installed, fail := isInstalled(toInstall.Name)

		if nil != fail {
			return fmt.Errorf(`%w;
error while checking whether or not '%s' package is installed;`, fail, toInstall.Name)
		}

		if !installed {
			result, fail := installPackage(toInstall.Name)

			if nil != fail {
				return fmt.Errorf(`%s;
%w;
error while installing '%s' package;`, result, fail, toInstall.Name)
			}
		}
	}

	return project.writeProject()
}

func (project *Project) AddToProject(packageName string) (fail error) {
	if project.isPackagePresent(packageName) {
		return fmt.Errorf("package '%s' already present in project", packageName)
	}

	show, fail := project.readPackage(packageName)

	if nil != fail {
		result, fail := installPackage(packageName)

		if nil != fail {
			return fmt.Errorf(`%w;
error while installing '%s' package;
due to: %s`, fail, packageName, result)
		}

		show, fail = packageShow(packageName)

		if nil != fail {
			return fmt.Errorf("%w;\nerror while fetching info from '%s' installed package", fail, packageName)
		}

		fail = project.addPackageToDatabase(show)

		if nil != fail {
			return fmt.Errorf("%w;\nerror while adding '%s' package info to database", fail, packageName)
		}
	}

	project.Packages = append(project.Packages, Package{
		Name:     show.Package,
		Revision: show.Revision,
	})

	return project.writeProject()
}

func (project *Project) RmFromProject(packageName string, uninstall bool) (fail error) {
	if !project.isPackagePresent(packageName) {
		return fmt.Errorf("package '%s' not present in project", packageName)
	}

	if uninstall {
		result, fail := uninstallPackage(packageName)

		if nil != fail {
			return fmt.Errorf(`%w;
error while uninstalling '%s' package;
due to: %s`, fail, packageName, result)
		}

		fail = project.rmPackage(Show{Package: packageName})

		if nil != fail {
			return fmt.Errorf("%w;\nerror while removing '%s' package from database", fail, packageName)
		}
	}

	fail = project.removePackage(packageName)

	if nil != fail {
		return fmt.Errorf("%w;\nerror while removing '%s' package", fail, packageName)
	}

	return fail
}

func (project *Project) UpgradeProjectPackage(projectPath string, packageName string) (fail error) {
	return fail
}

func ReadProject(projectPath string) (project Project, fail error) {
	project, fail = load(projectPath)

	if nil != fail {
		return project, fmt.Errorf("%w;\nerror while reading file from: %s", fail, projectPath)
	}

	fail = project.fetchDataBase()

	if nil != fail {
		return project, fmt.Errorf("%w;\nerror while reading Shojo's database: %s", fail, projectPath)
	}

	return project, fail
}
