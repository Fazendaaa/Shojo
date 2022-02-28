package controllers

import (
	"fmt"
	"time"

	shojo "github.com/Fazendaaa/Shojo/pkg"
	"github.com/theckman/yacspin"
)

func AddPackage(packages []string, path string) {
	config := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[35],
		Suffix:            " installing package",
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
		return
	}

	fail = spinner.Start()

	if nil != fail {
		return
	}

	for _, packageName := range packages {
		spinner.Message(fmt.Sprintf("'%s'", packageName))
		_, pkgFail := shojo.AddToProject(path, packageName)

		if nil != pkgFail {
			fmt.Printf(`%v;
error while installing package '%s';
halted execution
`, pkgFail, packageName)
			fail = spinner.StopFail()

			if nil != fail {
				return
			}

			return
		}
	}

	fail = spinner.Stop()

	if nil != fail {
		return
	}
}
