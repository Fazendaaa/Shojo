package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [projects' path]",
	Short: "Initialize's Shojo package manager",
	Long: `Initialize's Shojo package manager in the given directory creating
shojo.yml and feeding it with the needed packages set up`,
	Args: validPath,
	Run: func(cmd *cobra.Command, params []string) {
		controllers.InitProject(params[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
