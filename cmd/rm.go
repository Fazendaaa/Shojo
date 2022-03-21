package cmd

import (
	"fmt"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var (
	uninstall = false
	rmCmd     = &cobra.Command{
		Use:   "rm [package(s) to remove]",
		Short: "Removes a LaTex package to the current project",
		Long:  ``,
		Args:  validateProjectPath,
		Run: func(cmd *cobra.Command, params []string) {
			spinner, fail := createSpinner(" removing package", "")

			if nil != fail {
				fmt.Println()
				fmt.Println(fail)

				return
			}

			resultChannel, fail := controllers.RmPackages(params, projectPath, uninstall)

			if nil != fail {
				fmt.Println()
				fmt.Println(fail)

				return
			}

			fail = consumeChannel(params, "removing", spinner, resultChannel)

			if nil != fail {
				fmt.Println()
				fmt.Println(fail)

				killSpinner(spinner, false)

				return
			}

			killSpinner(spinner, true)
		},
	}
)

func init() {
	rmCmd.Flags().BoolVarP(&uninstall, "uninstall", "u", false, "optional flag to remove package not only from project but the machine as well")
	rmCmd.Flags().StringVarP(&projectPath, "project path", "p", ".", "optional shojo.yaml path to remove package from")
	rootCmd.AddCommand(rmCmd)
}
