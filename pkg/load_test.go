package shojo

import (
	"reflect"
	"testing"
)

func Testload(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		value, fail := load("../test/config/complete/")
		expected := Project{
			Tex: Tex{
				Version: "3.141592653",
			},
			TLMGR: TLMGR{
				Version: "61401",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("default", func(t *testing.T) {
		value, fail := load("../test/config/default/")
		expected := Project{
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("init", func(t *testing.T) {
		value, fail := load("../test/config/init/")
		expected := Project{
			Tex: Tex{
				Version: "3.141592653",
			},
			TLMGR: TLMGR{
				Version: "61401",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("missing", func(t *testing.T) {
		value, fail := load("../test/config/missing/")

		if nil == fail {
			t.Errorf("Expected to got a invalid missing project error but got the following value:\n%v", value)
		}
	})

	t.Run("named", func(t *testing.T) {
		value, fail := load("../test/config/named/foo.yml")
		expected := Project{
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("repository", func(t *testing.T) {
		value, fail := load("../test/config/repository/")
		expected := Project{
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("repository URL", func(t *testing.T) {
		value, fail := load("../test/config/repositoryURL/")
		expected := Project{
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tex", func(t *testing.T) {
		value, fail := load("../test/config/tex/")
		expected := Project{
			Tex: Tex{
				Version: "3.141592653",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tex Version", func(t *testing.T) {
		value, fail := load("../test/config/texVersion/")
		expected := Project{
			Tex: Tex{
				Version: "3.141592653",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tlmgr", func(t *testing.T) {
		value, fail := load("../test/config/tlmgr/")
		expected := Project{
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "61401",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tlmgr Version", func(t *testing.T) {
		value, fail := load("../test/config/tlmgrVersion/")
		expected := Project{
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "61401",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("yml", func(t *testing.T) {
		value, fail := load("../test/config/yml/")
		expected := Project{
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: []Package{
				{
					Name: "multirow",
				},
				{
					Name: "wrapfig",
				},
				{
					Name: "lastpage",
				},
				{
					Name: "hyphenat",
				},
				{
					Name: "hyphen-portuguese",
				},
				{
					Name: "babel-portuges",
				},
				{
					Name: "fancyhdr",
				},
				{
					Name: "tabu",
				},
				{
					Name: "varwidth",
				},
			},
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})
}
