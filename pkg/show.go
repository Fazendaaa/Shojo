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
error while generating regex for package 'show' command for '%s' package`,
			fail, packageName)
	}

	return regex.FindStringSubmatch(result)[1], fail
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

	show.CatVersion, fail = applyRegex("cat-version", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.CatLicense, fail = applyRegex("cat-license", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.CatTopics, fail = applyRegex("cat-topics", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.CatRelated, fail = applyRegex("cat-related", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	show.Collection, fail = applyRegex("collection", packageName, result)

	if nil != fail {
		return show, fmt.Errorf(`%w;
error while running regex process for '%s' package`, fail, packageName)
	}

	return show, fail
}
