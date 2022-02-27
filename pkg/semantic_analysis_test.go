package shojo

import (
	"testing"

	samael "github.com/Fazendaaa/Samael/pkg"
)

func TestSemanticAnalysis(t *testing.T) {
	t.Run("complete", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/complete/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("default", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/default/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("integer package", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/integerPackage/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil == fail {
			t.Errorf("Expect to get an error, but instead everything went ok given the following result:\n%v\n", value)
		}
	})

	t.Run("named", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/named/foo.yml", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("replicate repository", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/replicateRepository/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil == fail {
			t.Errorf("Expect to get an error, but instead everything went ok given the following result:\n%v\n", value)
		}
	})

	t.Run("repository", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/repository/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("repository URL", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/repositoryURL/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("tex", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/tex/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("tlmgr", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/tlmgr/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})

	t.Run("yml", func(t *testing.T) {
		lexed, _ := samael.LexProject("shojo", "../test/config/yml/", projectFunc)
		casted, _ := lexed.(Project)
		value, fail := semanticAnalysis(casted)

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})
}
