package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "shojo",
	Short: "Shojo is LaTex package manager",
	Long: `Shojo is made with Go to help you handle all of your LaTex package needs.
Complete documentation is available at https://github.com/Fazendaaa/Shojo

Shojo is also part of the Container tooling For Developers (CFD) initiative:
https://github.com/Fazendaaa/CFD`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
