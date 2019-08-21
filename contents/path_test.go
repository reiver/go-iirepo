package iirepo_contents_test

import (
	"github.com/reiver/go-iirepo/contents"

	"testing"
)

func TestPath(t *testing.T) {

	tests := []struct{
		RootPath string
		Expected string
	}{
		{
			RootPath: "/apple",
			Expected: "/apple/.ii/contents",
		},
		{
			RootPath: "/apple/BANANA",
			Expected: "/apple/BANANA/.ii/contents",
		},
		{
			RootPath: "/apple/BANANA/Cherry",
			Expected: "/apple/BANANA/Cherry/.ii/contents",
		},
		{
			RootPath: "/apple/BANANA/Cherry/dATE",
			Expected: "/apple/BANANA/Cherry/dATE/.ii/contents",
		},
	}

	for testNumber, test := range tests {
		actual := iirepo_contents.Path(test.RootPath)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual location is not what was expected.", testNumber)
			t.Logf("Root Path: %q", test.RootPath)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
