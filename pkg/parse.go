package shojo

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func checkURL(URL string) (fail error) {
	request, fail := url.ParseRequestURI(URL)

	if nil != fail {
		return fmt.Errorf(`URL presented isn't valid;
maybe missing "https://:" or even "https://:"?;
sometimes adding a trailing "/" might help you;
threw de following error %w:\n%s\n`, fail, request)
	}

	return fail
}

func checkTex(tex Tex) (_ Tex, fail error) {
	digits := 2
	version := strings.Split(tex.Version, ".")

	if 0 < len(tex.Version) && digits != len(version) {
		return tex, fmt.Errorf(`Tex version isn't valid, the presented version has
%d slices instead of %d`,
			len(version), digits)
	}

	return tex, fail
}

func checkTlmgr(tlmgr TLMGR) (_ TLMGR, fail error) {
	digits := 5
	version := numberOfCharacters(tlmgr.Version)

	if 0 < len(tlmgr.Version) && digits != version {
		return tlmgr, fmt.Errorf(`TLMGR version isn't valid, the presented version
has %d digits instead of %d`, version, digits)
	}

	return tlmgr, fail
}

// checkRepository https://stackoverflow.com/a/31480759/7092954
func checkRepository(repository Repository) (_ Repository, fail error) {
	if 0 == len(repository.URL) {
		return repository, fail
	}

	fail = checkURL(repository.URL)

	if nil != fail {
		return repository, fmt.Errorf("%w;\nURL presented in repository isn't valid.",
			fail)
	}

	return repository, fail
}

// checkPackages verify whether or not all packages are
// - without spaces
// - all lower case
// https://stackoverflow.com/a/14107337/7092954
// https://gobyexample.com/regular-expressions
func checkPackages(packages []Package) (_ []Package, fail error) {
	regex, _ := regexp.Compile("^[a-z_-]+$")

	for _, data := range packages {
		if !regex.MatchString(data.Name) {
			return packages, fmt.Errorf("package name '%s' is not a valid one", data.Name)
		}
	}

	return packages, fail
}

func parseProject(origin Project) (project Project, fail error) {
	project.filename = origin.filename
	project.Tex, fail = checkTex(origin.Tex)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid tex in project", fail)
	}

	project.TLMGR, fail = checkTlmgr(origin.TLMGR)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid TLMGR in project", fail)
	}

	project.Repository, fail = checkRepository(origin.Repository)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid repository in project", fail)
	}

	project.Packages, fail = checkPackages(origin.Packages)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid packages in project", fail)
	}

	return project, fail
}
