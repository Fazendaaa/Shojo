package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func validPath(cmd *cobra.Command, params []string) error {
	if 1 != len(params) {
		return fmt.Errorf(`missing project path!
If you want to start/install a project in the current path, you can run it again as:

	shojo init .
	shojo install .

Otherwise you may try to:

	shojo init /absolute/or/relative/project/path
	shojo install /absolute/or/relative/project/path`)
	}
	dir, fail := os.Stat(params[0])

	if fail != nil {
		return fmt.Errorf(`invalid path: %s
given error: %q`, params[0], fail)
	}
	if !dir.IsDir() {
		return fmt.Errorf("%q is not a directory", dir.Name())
	}

	return nil
}
