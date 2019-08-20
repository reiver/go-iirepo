package iirepo_stage_test

import (
	"github.com/reiver/go-iirepo/stage"

	"fmt"
)

func ExamplePath() {

	var parent string = "/home/joeblow/workspaces/myproject"

	repoStagePath := iirepo_stage.Path(parent)

	fmt.Printf("The repo's stage directory's path is: %s\n", repoStagePath)

	// Output:
	// The repo's stage directory's path is: /home/joeblow/workspaces/myproject/.ii/stage
}
