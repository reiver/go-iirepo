package iirepo_test

import (
	"github.com/reiver/go-iirepo"

	"testing"
)

func TestPath(t *testing.T) {

	tests := []struct{
		RootPath string
		Expected string
	}{
		{
			RootPath: "/apple",
			Expected: "/apple/.ii",
		},
		{
			RootPath: "/apple/BANANA",
			Expected: "/apple/BANANA/.ii",
		},
		{
			RootPath: "/apple/BANANA/Cherry",
			Expected: "/apple/BANANA/Cherry/.ii",
		},
		{
			RootPath: "/apple/BANANA/Cherry/dATE",
			Expected: "/apple/BANANA/Cherry/dATE/.ii",
		},



		{
			RootPath: "/home/joeblow/workspaces/golang/myproject",
			Expected: "/home/joeblow/workspaces/golang/myproject/.ii",
		},
	}

	for testNumber, test := range tests {
		actual := iirepo.Path(test.RootPath)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual location is not what was expected.", testNumber)
			t.Logf("Root Path: %q", test.RootPath)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
