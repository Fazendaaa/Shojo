package cmd

import (
	"fmt"
	"time"

	controllers "github.com/Fazendaaa/Shojo/controllers"
	"github.com/spf13/cobra"
	"github.com/theckman/yacspin"
)

var (
	uninstall = false
	rmCmd     = &cobra.Command{
		Use:   "rm [package(s) to remove]",
		Short: "Removes a LaTex package to the current project",
		Long:  ``,
		Args:  validateProjectPath,
		Run: func(cmd *cobra.Command, params []string) {
			config := yacspin.Config{
				Frequency:         100 * time.Millisecond,
				CharSet:           yacspin.CharSets[35],
				Suffix:            " removing package",
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
				fail = controllers.RmPackages(packageName, projectPath, uninstall)

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
)

func init() {
	rmCmd.Flags().BoolVarP(&uninstall, "uninstall", "u", false, "optional flag to remove package not only from project but the machine as well")
	rmCmd.Flags().StringVarP(&projectPath, "project path", "p", ".", "optional shojo.yaml path to remove package from")
	rootCmd.AddCommand(rmCmd)
}
