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

	t.Run("tex", func(t *testing.T) {
		value, fail := samael.LexProject("shojo", "../test/config/tex/", projectFunc)

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
