package shojo

import "regexp"

func tlmgr(subcommand string, paramters []string) (result string, fail error) {
	return shell("tlmgr", append([]string{subcommand}, paramters...))
}

func install(packageName string) (fail error) {
	_, fail = tlmgr("install", []string{packageName})

	return fail
}

func tlmgrVersion() (version string, fail error) {
	version, fail = tlmgr("--version", []string{})

	if nil != fail {
		return version, fail
	}

	regex, fail := regexp.Compile(`^tlmgr\s*revision\s*(\d*)`)

	if nil != fail {
		return version, fail
	}

	return regex.FindStringSubmatch(version)[1], fail
}
