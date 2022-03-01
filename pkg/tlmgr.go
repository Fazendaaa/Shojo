package shojo

import (
	"fmt"
	"regexp"
)

type TLMGR struct {
	Version string `yaml:"version"`
}

func tlmgr(subcommand string, paramters []string) (result string, fail error) {
	return shell("tlmgr", append([]string{subcommand}, paramters...))
}

func installPackage(packageName string) (result string, fail error) {
	return tlmgr("install", []string{packageName})
}

func uninstallPackage(packageName string) (result string, fail error) {
	return tlmgr("remove", []string{"--force", packageName})
}

func tlmgrVersion() (version string, fail error) {
	version, fail = tlmgr("--version", []string{})

	if nil != fail {
		return version, fmt.Errorf("%w;\nerror while fecthing tlmgr version", fail)
	}

	regex, fail := regexp.Compile(`^tlmgr\s*revision\s*(\d*)`)

	if nil != fail {
		return version, fmt.Errorf("%w;\nerror while parsing tlmgr version", fail)
	}

	return regex.FindStringSubmatch(version)[1], fail
}

func isInstalled(packageName string) (ok bool, fail error) {
	result, fail := tlmgr("info", append([]string{"--only-installed"}))

	if nil != fail {
		return ok, fmt.Errorf("%w;\nerror while fecthing tlmgr version", fail)
	}

	ok, fail = regexp.MatchString(fmt.Sprintf(`i\s*%s:`, packageName), result)

	if nil != fail {
		return ok, fmt.Errorf("%w;\nerror while parsing tlmgr version", fail)
	}

	return ok, fail
}
