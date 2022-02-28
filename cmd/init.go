package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize's Shojo package manager",
	Long: `Initialize's Shojo package manager in the given direcotry creating
shojo.yml and feeding it with the needed packages set up`,
	Run: func(cmd *cobra.Command, params []string) {
		controllers.InitProject()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
