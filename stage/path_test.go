package iirepo_stage_test

import (
	"github.com/reiver/go-iirepo/stage"

	"testing"
)

func TestPath(t *testing.T) {

	tests := []struct{
		RootPath string
		Expected string
	}{
		{
			RootPath: "/apple",
			Expected: "/apple/.ii/stage",
		},
		{
			RootPath: "/apple/BANANA",
			Expected: "/apple/BANANA/.ii/stage",
		},
		{
			RootPath: "/apple/BANANA/Cherry",
			Expected: "/apple/BANANA/Cherry/.ii/stage",
		},
		{
			RootPath: "/apple/BANANA/Cherry/dATE",
			Expected: "/apple/BANANA/Cherry/dATE/.ii/stage",
		},
	}

	for testNumber, test := range tests {
		actual := iirepo_stage.Path(test.RootPath)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual location is not what was expected.", testNumber)
			t.Logf("Root Path: %q", test.RootPath)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
