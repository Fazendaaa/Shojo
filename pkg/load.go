package shojo

import (
	samael "github.com/Fazendaaa/Samael/pkg"
)

// Load
func Load(projectPath string) (project Project, fail error) {
	lexed, fail := samael.LexProject("shojo", projectPath, projectFunc)

	if nil != fail {
		return project, fail
	}

	casted, ok := lexed.(Project)

	if !ok {
		return project, fail
	}

	parsed, fail := parseProject(casted)

	if nil != fail {
		return parsed, fail
	}

	return semanticAnalysis(parsed)
}
