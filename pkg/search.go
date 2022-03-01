package shojo

import "regexp"

func packageSearch(packageName string) (version string, fail error) {
	version, fail = tlmgr("search", []string{"--global", "--all", packageName})

	if nil != fail {
		return version, fail
	}

	regex, fail := regexp.Compile(`^tlmgr\s*revision\s*(\d*)`)

	if nil != fail {
		return version, fail
	}

	return regex.FindStringSubmatch(version)[1], fail
}
