package shojo

import (
	"fmt"
	"regexp"
)

type Show struct {
	Package     string `yaml:"package"`
	Category    string `yaml:"category"`
	Shortdesc   string `yaml:"shortdesc"`
	Longdesc    string `yaml:"longdesc"`
	Installed   string `yaml:"installed"`
	Revision    string `yaml:"revision"`
	Sizes       string `yaml:"sizes"`
	Relocatable string `yaml:"relocatable"`
	CatVersion  string `yaml:"cat-version"`
	CatLicense  string `yaml:"cat-license"`
	CatTopics   string `yaml:"cat-topics"`
	CatRelated  string `yaml:"cat-related"`
	Collection  string `yaml:"collection"`
}

func applyRegex(topic string, packageName string, show string) (result string, fail error) {
	regex, fail := regexp.Compile(`^` + packageName + `:\s*(?.*)$`)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while generating regex for package 'show' command for '%s' package`, fail, packageName)
	}

	return regex.FindStringSubmatch(result)[1], fail
}

func packageShow(packageName string) (show Show, fail error) {
	result, fail := tlmgr("show", []string{packageName})

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running 'show' command for '%s' package`, fail, packageName)
	}

	simpleFields := []string{"package", "category", "shortdesc", "longdesc",
		"installed", "sizes", "relocatable", "collection"}
	composedFields := []string{"cat-version", "cat-license", "cat-topics", "cat-related"}

	show.Package, fail = applyRegex("package", packageName, result)

	return show, fail
}
