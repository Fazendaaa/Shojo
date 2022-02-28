package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var projectPath string
var addCmd = &cobra.Command{
	Use:   "add [package(s) to add]",
	Short: "Adds a LaTex package to the current project",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, params []string) {
		controllers.AddPackage(params)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional shojo.yaml path to add package to")
}
