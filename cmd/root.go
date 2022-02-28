package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shojo",
	Short: "Shojo is LaTex package manager",
	Long: `Shojo is made with Go to help you handle all of your LaTex package
needs. Shojo is also part of the Container For Developers (CFD) initiative.
Complete documentation is available at https://github.com/Fazendaaa/Shojo`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
