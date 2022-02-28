package shojo

import (
	"regexp"
)

type Tex struct {
	Version string `yaml:"version"`
}

func tex(subcommand string, paramters []string) (result string, fail error) {
	return shell("tex", append([]string{subcommand}, paramters...))
}

func texVersion() (version string, fail error) {
	version, fail = tex("--version", []string{})

	if nil != fail {
		return version, fail
	}

	regex, fail := regexp.Compile(`^TeX\s*(\d*.\d*)`)

	if nil != fail {
		return version, fail
	}

	return regex.FindStringSubmatch(version)[1], fail
}
