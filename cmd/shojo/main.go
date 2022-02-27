package main

import (
	"github.com/spf13/cobra"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "shojo",
	}

	for _, command := range shojo.ShojoCMD {
		translated := &cobra.Command{
			Use:   command.Name,
			Short: command.Usage.Short,
			Long:  command.Usage.Long,
			Args:  cobra.MinimumNArgs(len(command.Params)),
			Run: func(cmd *cobra.Command, args []string) {
				command.Function(args)
			},
		}

		rootCmd.AddCommand(translated)
	}

	rootCmd.Execute()
}
