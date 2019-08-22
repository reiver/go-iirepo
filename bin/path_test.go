package iirepo_bin_test

import (
	"github.com/reiver/go-iirepo/bin"

	"testing"
)

func TestPath(t *testing.T) {

	tests := []struct{
		RootPath string
		Expected string
	}{
		{
			RootPath: "/apple",
			Expected: "/apple/.ii/bin",
		},
		{
			RootPath: "/apple/BANANA",
			Expected: "/apple/BANANA/.ii/bin",
		},
		{
			RootPath: "/apple/BANANA/Cherry",
			Expected: "/apple/BANANA/Cherry/.ii/bin",
		},
		{
			RootPath: "/apple/BANANA/Cherry/dATE",
			Expected: "/apple/BANANA/Cherry/dATE/.ii/bin",
		},
	}

	for testNumber, test := range tests {
		actual := iirepo_bin.Path(test.RootPath)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual location is not what was expected.", testNumber)
			t.Logf("Root Path: %q", test.RootPath)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
