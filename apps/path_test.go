package iirepo_apps_test

import (
	"github.com/reiver/go-iirepo/apps"

	"testing"
)

func TestPath(t *testing.T) {

	tests := []struct{
		RootPath string
		Expected string
	}{
		{
			RootPath: "/apple",
			Expected: "/apple/.ii/apps",
		},
		{
			RootPath: "/apple/BANANA",
			Expected: "/apple/BANANA/.ii/apps",
		},
		{
			RootPath: "/apple/BANANA/Cherry",
			Expected: "/apple/BANANA/Cherry/.ii/apps",
		},
		{
			RootPath: "/apple/BANANA/Cherry/dATE",
			Expected: "/apple/BANANA/Cherry/dATE/.ii/apps",
		},
	}

	for testNumber, test := range tests {
		actual := iirepo_apps.Path(test.RootPath)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual location is not what was expected.", testNumber)
			t.Logf("Root Path: %q", test.RootPath)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
