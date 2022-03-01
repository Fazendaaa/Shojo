package cmd

import (
	"fmt"

	"github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo [new repository to use]",
	Short: "Configures the repository from where to fetch packages from",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := createSpinner(" setting repository", "")

		if nil != fail {
			fmt.Printf("\n%v", fail)

			return
		}

		fail = controllers.SetRepository(params[0], projectPath)

		if nil != fail {
			fmt.Printf("\n%v", fail)

			killSpinner(spinner, false)

			return
		}

		killSpinner(spinner, true)
	},
}

func init() {
	repoCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional shojo.yaml path to change repository")
	rootCmd.AddCommand(repoCmd)
}
