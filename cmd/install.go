package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [projects' path]",
	Short: "Install all presented LaTex packages in the current project",
	Args:  validPath("install"),
	Run: func(cmd *cobra.Command, params []string) {
		controllers.InstallProject()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
