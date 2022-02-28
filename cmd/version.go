package cmd

import (
	"fmt"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints Shojo's version",
	Long:  `All software has versions. This is Shojo's`,
	Run: func(cmd *cobra.Command, params []string) {
		fmt.Println(controllers.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
