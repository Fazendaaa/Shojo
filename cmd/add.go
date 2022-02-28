package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a LaTex package to the current project",
	Long:  ``,
	Run: func(cmd *cobra.Command, params []string) {
		controllers.AddPackage(params)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
