package cmd

import (
	"fmt"
	"os"

	shojo "github.com/Fazendaaa/Shojo/pkg"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a LaTex package to the current project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, fail := os.Getwd()

		if nil != fail {
			fmt.Errorf("%w;\ncould not read current directory", fail)
		}

		project, fail := shojo.Load(path)

		fmt.Println(project)

		if nil != fail {
			fmt.Errorf("%w;\nmalformed tex definition", fail)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
