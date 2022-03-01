package cmd

import (
	"fmt"
	"time"

	"github.com/theckman/yacspin"
)

func createSpinner(suffix string, message string) (spinner *yacspin.Spinner, fail error) {
	config := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[35],
		Suffix:            suffix,
		SuffixAutoColon:   true,
		Message:           message,
		StopCharacter:     "✓",
		StopMessage:       "done",
		StopFailCharacter: "✗",
		StopFailMessage:   "failed",
		StopColors:        []string{"fgGreen"},
	}
	spinner, fail = yacspin.New(config)

	if nil != fail {
		return spinner, fmt.Errorf("\n%v", fail)
	}

	fail = spinner.Start()

	if nil != fail {
		return spinner, fmt.Errorf("\n%v", fail)
	}

	return spinner, fail
}

func killSpinner(spinner *yacspin.Spinner, ok bool) (fail error) {
	if !ok {
		fail = spinner.StopFail()
		if nil != fail {
			fmt.Printf("\n%v", fail)
		}
	}

	fail = spinner.Stop()

	if nil != fail {
		fmt.Printf("\n%v", fail)
	}

	return fail
}
