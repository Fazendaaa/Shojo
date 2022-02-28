package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Shojo",
	Long:  `All software has versions. This is Shojo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Shojo is LaTex package manager v0.0 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
