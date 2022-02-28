package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [projects' path]",
	Short: "Removes a LaTex package to the current project",
	Long:  ``,
	Args:  validateProjectPath,
	Run: func(cmd *cobra.Command, params []string) {
		controllers.RmPackages(params, projectPath)
	},
}

func init() {
	rmCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional shojo.yaml path to remove package from")
	rootCmd.AddCommand(rmCmd)
}
