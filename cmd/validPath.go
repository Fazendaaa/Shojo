package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func validPath(command string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, params []string) error {
		if 1 != len(params) {
			return fmt.Errorf(`missing project path!
If you want to %s a project in the current path, you can run it again as:

	shojo %s .

Otherwise you may try to:

	shojo %s /absolute/or/relative/project/path`, command, command, command)
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
}
