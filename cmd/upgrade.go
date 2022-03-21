package cmd

import (
	"fmt"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade [projects' path]",
	Short: "Upgrades all LaTex packages in the current project",
	Long:  ``,
	Args:  validPath("upgrade"),
	Run: func(cmd *cobra.Command, params []string) {
		spinner, fail := createSpinner(" upgrading packages", "")

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		resultChannel, fail := controllers.UpgradePackages(projectPath)

		if nil != fail {
			fmt.Println()
			fmt.Println(fail)

			return
		}

		fail = consumeChannel(params, "upgrading", spinner, resultChannel)

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
	upgradeCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional shojo.yaml path to add package to")
	rootCmd.AddCommand(upgradeCmd)
}
