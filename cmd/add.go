package cmd

import (
	"fmt"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [package(s) to add]",
	Short: "Adds a LaTex package to the current project",
	Long:  ``,
	Args:  validateProjectPath,
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := createSpinner(" adding package", "")

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		resultChannel, fail := controllers.AddPackages(params, projectPath)

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		fail = consumeChannel(params, "installing", spinner, resultChannel)

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			killSpinner(spinner, false)

			return
		}

		killSpinner(spinner, true)
	},
}

func init() {
	addCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional shojo.yaml path to add package to")
	rootCmd.AddCommand(addCmd)
}
