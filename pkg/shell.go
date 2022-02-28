package shojo

import (
	"fmt"
	"os/exec"
)

// shell Handles all terminal realted calls
// https://stackoverflow.com/a/7786922/7092954
func shell(command string, parameters []string) (result string, fail error) {
	cmd := exec.Command(command, parameters...)
	stdout, fail := cmd.Output()

	if nil != fail {
		return result, fmt.Errorf(`%w;
error while executing command: %s %+q`, fail, command, parameters)
	}

	return string(stdout), fail
}
