package shojo

import (
	"testing"

	samael "github.com/Fazendaaa/Samael/pkg"
)

func TestParse(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/complete/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("default", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/default/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("integer repository", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/integerRepository/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("integer Tex Version", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/integerTexVersion/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("invalid Repository", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/invalidRepository/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil == fail {
			t.Errorf("Expected to got a invalid URL error but got the following value:\n%v", value)
		}
	})

	t.Run("named", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/named/foo.yml", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("repository", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/repository/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("repository URL", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/repositoryURL/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("string Tlmgr Version", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/stringTlmgrVersion/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("tex", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/tex/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("tlmgr", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/tlmgr/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("yml", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/yml/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := parseProject(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})
}
