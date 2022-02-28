package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install all presented LaTex packages in the current project",
	Long:  ``,
	Run: func(cmd *cobra.Command, params []string) {
		controllers.InstallProject()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
