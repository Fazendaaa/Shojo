package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var projectPath string

func checkPath(path string) error {
	dir, fail := os.Stat(path)

	if nil != fail {
		return fmt.Errorf(`invalid path: %s
	given error: %q`, path, fail)
	}
	if !dir.IsDir() {
		return fmt.Errorf("%q is not a directory", dir.Name())
	}

	return nil
}

func validPath(command string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, params []string) error {
		if 1 != len(params) {
			return fmt.Errorf(`missing project path!
If you want to %s a project in the current path, you can run it again as:

	shojo %s .

Otherwise you may try to:

	shojo %s /absolute/or/relative/project/path`, command, command, command)
		}

		return checkPath(params[0])
	}
}

func validateProjectPath(cmd *cobra.Command, params []string) error {
	if 1 > len(params) {
		return fmt.Errorf("requires at least 1 arg(s), only received 0")
	}

	if "" == projectPath {
		path, fail := os.Getwd()

		if nil != fail {
			fmt.Printf("%v;\ncould not read current directory", fail)
		}

		projectPath = path
	}

	return checkPath(projectPath)
}
