package cmd

import (
	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a LaTex package to the current project",
	Long:  ``,
	Run: func(cmd *cobra.Command, params []string) {
		controllers.RmPackages(params)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
