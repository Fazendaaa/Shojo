package shojo

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		value, fail := Load("../test/config/complete/")
		expected := Project{
			tex: Tex{
				version: "6.3.3",
			},
			tlmgr: TLMGR{
				version: "61401",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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
		value, fail := Load("../test/config/default/")
		expected := Project{
			tex: Tex{
				version: "",
			},
			tlmgr: TLMGR{
				version: "",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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

	t.Run("named", func(t *testing.T) {
		value, fail := Load("../test/config/named/foo.yml")
		expected := Project{
			tex: Tex{
				version: "",
			},
			tlmgr: TLMGR{
				version: "",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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
		value, fail := Load("../test/config/repository/")
		expected := Project{
			tex: Tex{
				version: "",
			},
			tlmgr: TLMGR{
				version: "",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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
		value, fail := Load("../test/config/repositoryURL/")
		expected := Project{
			tex: Tex{
				version: "",
			},
			tlmgr: TLMGR{
				version: "",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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
		value, fail := Load("../test/config/tex/")
		expected := Project{
			tex: Tex{
				version: "6.3.3",
			},
			tlmgr: TLMGR{
				version: "",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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
		value, fail := Load("../test/config/tlmgr/")
		expected := Project{
			tex: Tex{
				version: "",
			},
			tlmgr: TLMGR{
				version: "61401",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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
		value, fail := Load("../test/config/yml/")
		expected := Project{
			tex: Tex{
				version: "",
			},
			tlmgr: TLMGR{
				version: "",
			},
			repository: Repository{
				url: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			packages: []Package{
				{
					name: "multirow",
				},
				{
					name: "wrapfig",
				},
				{
					name: "lastpage",
				},
				{
					name: "hyphenat",
				},
				{
					name: "hyphen-portuguese",
				},
				{
					name: "babel-portuges",
				},
				{
					name: "fancyhdr",
				},
				{
					name: "tabu",
				},
				{
					name: "varwidth",
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
