package iirepo_stage_test

import (
	"github.com/reiver/go-iirepo/stage"

	"fmt"
)

func ExampleName() {

	repoStageName := iirepo_stage.Name()

	fmt.Printf("The repo stage directory's name is: %s\n", repoStageName)

	// Output:
	// The repo stage directory's name is: stage
}
