package shojo

import (
	"fmt"

	samael "github.com/Fazendaaa/Samael/pkg"
)

// load
func load(projectPath string) (project Project, fail error) {
	lexed, fail := samael.LexProject("shojo", projectPath, projectFunc)

	if nil != fail {
		return project, fail
	}

	casted, ok := lexed.(Project)

	if !ok {
		return project, fmt.Errorf("cannot normalize project")
	}

	parsed, fail := parseProject(casted)

	if nil != fail {
		return parsed, fail
	}

	return semanticAnalysis(parsed)
}
