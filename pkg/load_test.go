package shojo

import "testing"

func TestLoad(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		value, fail := Load("../test/config/complete/")

		if nil != fail {
			t.Errorf("got:\n%v\n and the given error condition is:\n%s", value, fail)
		}
	})
}
