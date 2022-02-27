package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	shojo "github.com/Fazendaaa/Shojo/pkg"
)

func main() {
	shojoCMD := make([](shojo.CMD), 2)
	rootCmd := &cobra.Command{
		Use: "shojo",
	}

	for _, command := range shojoCMD {
		translated := &cobra.Command{
			Use:   command.Name,
			Short: command.Usage.Short,
			Long:  command.Usage.Long,
			Args:  cobra.MinimumNArgs(len(command.Params)),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Print: " + strings.Join(args, " "))
				command.Function(args)
			},
		}

		rootCmd.AddCommand(translated)
	}

	rootCmd.Execute()
}
