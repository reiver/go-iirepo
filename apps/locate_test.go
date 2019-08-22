package iirepo_app

import (
	"testing"
)

func TestLocate(t *testing.T) {

	tests := []struct{
		Path     string
		RepoPath string
		Expected string
	}{
		{
			Path:     "/apple/banana/cherry/date/file.txt",
			RepoPath: "/apple",
			Expected: "/apple/.ii/app",
		},
		{
			Path:     "/apple/banana/cherry/date/file.txt",
			RepoPath: "/apple/banana",
			Expected: "/apple/banana/.ii/app",
		},
		{
			Path:     "/apple/banana/cherry/date/file.txt",
			RepoPath: "/apple/banana/cherry",
			Expected: "/apple/banana/cherry/.ii/app",
		},
	}

	for testNumber, test := range tests {

		fn := func(string)(string, error) {
			return test.RepoPath, nil
		}

		contentspath, err := locate(test.Path, fn)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, contentspath; expected != actual {
			t.Errorf("For test #%d, the contents path that was actually gotten was not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}