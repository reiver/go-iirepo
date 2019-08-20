package iirepo_test

import (
	"github.com/reiver/go-iirepo"

	"fmt"
)

func ExamplePath() {

	var parent string = "/home/joeblow/workspaces/myproject"

	repoPath := iirepo.Path(parent)

	fmt.Printf("The repo directory's path is: %s\n", repoPath)

	// Output:
	// The repo directory's path is: /home/joeblow/workspaces/myproject/.ii
}
