package shojo

import "fmt"

type Repository struct {
	URL string `yaml:"url"`
}

func SetRepository(repository string, projectPath string) (fail error) {
	fail = checkURL(repository)

	if nil != fail {
		return fmt.Errorf("%w;\nerror while checking the given repository URL", fail)
	}

	project, fail := load(projectPath)

	if nil != fail {
		return fmt.Errorf("%w\n; error while loading project definition while configuring repository", fail)
	}

	project.Repository.URL = repository

	return writeToProject(project)
}
