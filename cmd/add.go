package cmd

import (
	"fmt"
	"time"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
	"github.com/theckman/yacspin"
)

var addCmd = &cobra.Command{
	Use:   "add [package(s) to add]",
	Short: "Adds a LaTex package to the current project",
	Long:  ``,
	Args:  validateProjectPath,
	Run: func(cmd *cobra.Command, params []string) {
		config := yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[35],
			Suffix:            " adding package",
			SuffixAutoColon:   true,
			Message:           "exporting data",
			StopCharacter:     "✓",
			StopMessage:       "done",
			StopFailCharacter: "✗",
			StopFailMessage:   "failed",
			StopColors:        []string{"fgGreen"},
		}
		spinner, fail := yacspin.New(config)

		if nil != fail {
			fmt.Printf("\n%v", fail)
			return
		}

		fail = spinner.Start()

		if nil != fail {
			fmt.Printf("\n%v", fail)
			return
		}

		for _, packageName := range params {
			spinner.Message(fmt.Sprintf("'%s'", packageName))
			fail = controllers.AddPackage(packageName, projectPath)

			if nil != fail {
				fmt.Printf("\n%v", fail)
				fail = spinner.StopFail()

				if nil != fail {
					fmt.Printf("\n%v", fail)
				}

				return
			}
		}

		fail = spinner.Stop()

		if nil != fail {
			fmt.Printf("\n%v", fail)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&projectPath, "project path", "p", "", "optional shojo.yaml path to add package to")
	rootCmd.AddCommand(addCmd)
}
