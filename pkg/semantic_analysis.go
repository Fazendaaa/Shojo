package shojo

import "fmt"

func semanticTex(tex Tex) (_ Tex, fail error) {
	return tex, fail
}

func semanticTlmgr(tlmgr TLMGR) (_ TLMGR, fail error) {
	return tlmgr, fail
}

func semanticRepository(repository Repository) (_ Repository, fail error) {
	if 0 == len(repository.URL) {
		return DEFAULT_REPOSITORY, fail
	}

	return repository, fail
}

func semanticPackages(packages []Package) (_ []Package, fail error) {
	uniquePackages := unique(packagesToInterface(packages))

	if len(packages) != len(uniquePackages) {
		return packages, fmt.Errorf("packages has duplicate entries")
	}

	return packages, fail
}

// semanticAnalysis goes trough a project definition and checks for:
// - ambiguity
// - adds missing variables not set by the user in their config file
// It returns the sanitized version of the Project or the error related to it
func semanticAnalysis(origin Project) (project Project, fail error) {
	project.filename = origin.filename
	project.Tex, fail = semanticTex(origin.Tex)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid tex semantic in project", fail)
	}

	project.TLMGR, fail = semanticTlmgr(origin.TLMGR)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid tlmgr semantic in project", fail)
	}

	project.Repository, fail = semanticRepository(origin.Repository)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid repository semantic in project", fail)
	}

	project.Packages, fail = semanticPackages(origin.Packages)

	if nil != fail {
		return project, fmt.Errorf("%w;\ninvalid packages semantic in project", fail)
	}

	return project, fail
}
