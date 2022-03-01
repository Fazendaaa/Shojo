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
	pattern := fmt.Sprintf(`%s:\s*(.*)\n`, topic)
	regex, fail := regexp.Compile(pattern)

	if nil != fail {
		return result, fmt.Errorf(`%w;
error while generating regex for 'show' command for '%s' package in '%s' topic`,
			fail, packageName, topic)
	}

	matches := regex.FindStringSubmatch(show)

	if 1 > len(matches) {
		return show, fmt.Errorf(`%w;
error while processing information for 'show' command for '%s' package in '%s' topic`,
			fail, packageName, topic)
	}

	return matches[1], fail
}

func packageShow(packageName string) (show Show, fail error) {
	result, fail := tlmgr("show", []string{packageName})

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running 'show' command for '%s' package`, fail, packageName)
	}

	show.Package, fail = applyRegex("package", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Category, fail = applyRegex("category", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Shortdesc, fail = applyRegex("shortdesc", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Longdesc, fail = applyRegex("longdesc", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Installed, fail = applyRegex("installed", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Revision, fail = applyRegex("revision", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Sizes, fail = applyRegex("sizes", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Relocatable, fail = applyRegex("relocatable", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.CatVersion, _ = applyRegex("cat-version", packageName, result)
	show.CatLicense, _ = applyRegex("cat-license", packageName, result)
	show.CatTopics, _ = applyRegex("cat-topics", packageName, result)
	show.CatRelated, _ = applyRegex("cat-related", packageName, result)
	show.Collection, _ = applyRegex("collection", packageName, result)

	return show, fail
}
