package shojo

import (
	"path/filepath"
	"reflect"
	"testing"
)

var (
	PACKAGES = []Package{
		{
			Name:     "multirow",
			Revision: "58396",
		},
		{
			Name:     "wrapfig",
			Revision: "61719",
		},
		{
			Name:     "lastpage",
			Revision: "60414",
		},
		{
			Name:     "hyphenat",
			Revision: "15878",
		},
		{
			Name:     "hyphen-portuguese",
			Revision: "58609",
		},
		{
			Name:     "babel-portuges",
			Revision: "59883",
		},
		{
			Name:     "fancyhdr",
			Revision: "57672",
		},
		{
			Name:     "tabu",
			Revision: "61719",
		},
		{
			Name:     "varwidth",
			Revision: "24104",
		},
	}
)

func TestLoad(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		path := "../test/config/complete/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "3.141592653",
			},
			TLMGR: TLMGR{
				Version: "61401",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("default", func(t *testing.T) {
		path := "../test/config/default/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("init", func(t *testing.T) {
		path := "../test/config/init/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
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
		path := "../test/config/missing/"
		value, fail := load(path)

		if nil == fail {
			t.Errorf("Expected to got a invalid missing project error but got the following value:\n%v", value)
		}
	})

	t.Run("named", func(t *testing.T) {
		path := "../test/config/named/"
		file := "foo.yml"
		filename := filepath.Join(path, file)
		value, fail := load(filename)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("repository", func(t *testing.T) {
		path := "../test/config/repository/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("repository URL", func(t *testing.T) {
		path := "../test/config/repositoryURL/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tex", func(t *testing.T) {
		path := "../test/config/tex/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "3.141592653",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tex Version", func(t *testing.T) {
		path := "../test/config/texVersion/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "3.141592653",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tlmgr", func(t *testing.T) {
		path := "../test/config/tlmgr/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "61401",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("tlmgr Version", func(t *testing.T) {
		path := "../test/config/tlmgrVersion/"
		file := "shojo.yaml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "61401",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})

	t.Run("yml", func(t *testing.T) {
		path := "../test/config/yml/"
		file := "shojo.yml"
		filename := filepath.Join(path, file)
		value, fail := load(path)
		expected := Project{
			filename: filename,
			Tex: Tex{
				Version: "",
			},
			TLMGR: TLMGR{
				Version: "",
			},
			Repository: Repository{
				URL: "https://mirror.ctan.org/systems/texlive/tlnet/",
			},
			Packages: PACKAGES,
		}

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}

		if !reflect.DeepEqual(expected, value) {
			t.Errorf("got mismatching configurations:\n%#v\n\n%#v", expected, value)
		}
	})
}
