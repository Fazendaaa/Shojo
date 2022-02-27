package shojo

import (
	"os/exec"
)

// shell Handles all terminal realted calls
// https://stackoverflow.com/a/7786922/7092954
func shell(command string, parameters []string) (result string, fail error) {
	cmd := exec.Command(command, parameters...)
	stdout, err := cmd.Output()

	if err != nil {
		return result, err
	}

	return string(stdout), nil
}
