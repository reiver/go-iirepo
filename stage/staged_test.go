package iirepo_stage

import (
	"testing"
)

func TestStaged(t *testing.T) {

	tests := []struct{
		SrcPath  string
		RootPath string
		StagePath string
		Expected string
	}{
		{
			SrcPath:   "/apple/banana/cherry/date/file.txt",
			RootPath:  "/apple",
			StagePath: "/apple/.ii/stage",
			Expected:  "/apple/.ii/stage/banana/cherry/date/file.txt",
		},
		{
			SrcPath:   "/apple/banana/cherry/date/file.txt",
			RootPath:  "/apple/banana",
			StagePath: "/apple/banana/.ii/stage",
			Expected:  "/apple/banana/.ii/stage/cherry/date/file.txt",
		},
		{
			SrcPath:   "/apple/banana/cherry/date/file.txt",
			RootPath:  "/apple/banana/cherry",
			StagePath: "/apple/banana/cherry/.ii/stage",
			Expected:  "/apple/banana/cherry/.ii/stage/date/file.txt",
		},
	}

	for testNumber, test := range tests {

		stagepath, err := stagedPath(test.SrcPath, test.RootPath, test.StagePath)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, stagepath; expected != actual {
			t.Errorf("For test #%d, the stage path that was actually gotten was not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
