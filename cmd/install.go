package cmd

import (
	"fmt"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [projects' path]",
	Short: "Install all presented LaTex packages in the current project",
	Args:  validPath("install"),
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := createSpinner(" installing packages", "")

		if nil != fail {
			fmt.Printf("\n%v", fail)

			return
		}

		fail = controllers.InstallProject(params[0])

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
	rootCmd.AddCommand(installCmd)
}
