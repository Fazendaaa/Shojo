package cmd

import (
	"fmt"
	"os"

	shojo "github.com/Fazendaaa/Shojo/pkg"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a LaTex package to the current project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, fail := os.Getwd()

		if nil != fail {
			fmt.Errorf("%w;\ncould not read current directory", fail)
		}

		_, fail = shojo.Load(path)

		if nil != fail {
			fmt.Errorf("%w;\nmalformed tex definition", fail)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
