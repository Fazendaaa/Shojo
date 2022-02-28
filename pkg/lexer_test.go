package shojo

import (
	"testing"

	samael "github.com/Fazendaaa/Samael/pkg"
)

func TestLex(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/complete/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("default", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/default/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("init", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/init/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("integer package", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/integerPackage/", projectFunc)

		if nil == fail {
			t.Errorf("Expected to get an error but got the response: \n%v", value)
		}
	})

	t.Run("integer repository", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/integerRepository/", projectFunc)

		if nil == fail {
			t.Errorf("Expected to get an error but got the response: \n%v", value)
		}
	})

	t.Run("integer Tex Version", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/integerTexVersion/", projectFunc)

		if nil == fail {
			t.Errorf("Expected to get an error but got the response: \n%v", value)
		}
	})

	t.Run("missing packages", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/missingPackages/", projectFunc)

		if nil == fail {
			t.Errorf("Expected to get an error but got the response: \n%v", value)
		}
	})

	t.Run("named", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/named/foo.yml", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("repository", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/repository/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("repository URL", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/repositoryURL/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("string Tlmgr Version", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/stringTlmgrVersion/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("tex", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/tex/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("tex Version", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/texVersion/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("tlmgr", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/tlmgr/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("yml", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/yml/", projectFunc)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})
}
