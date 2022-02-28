package cmd

import (
	"fmt"
	"os"

	shojo "github.com/Fazendaaa/Shojo/pkg"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize's Shojo package manager",
	Long: `Initialize's Shojo package manager in the given direcotry creating
shojo.yml and feeding it with the needed packages set up`,
	Run: func(cmd *cobra.Command, args []string) {
		path, fail := os.Getwd()

		if nil != fail {
			fmt.Errorf("%w;\ncould not read current directory", fail)
		}

		shojo.CreateProject(path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
