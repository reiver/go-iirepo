package iirepo_stage_test

import (
	"github.com/reiver/go-iirepo/stage"

	"testing"
)

func TestPath(t *testing.T) {

	tests := []struct{
		Path string
		Expected string
	}{
		{
			Path: "/apple",
			Expected: "/apple/.ii/stage",
		},
		{
			Path: "/apple/BANANA",
			Expected: "/apple/BANANA/.ii/stage",
		},
		{
			Path: "/apple/BANANA/Cherry",
			Expected: "/apple/BANANA/Cherry/.ii/stage",
		},
		{
			Path: "/apple/BANANA/Cherry/dATE",
			Expected: "/apple/BANANA/Cherry/dATE/.ii/stage",
		},
	}

	for testNumber, test := range tests {
		actual := iirepo_stage.Path(test.Path)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, the actual location is not what was expected.", testNumber)
			t.Logf("Path: %q", test.Path)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
