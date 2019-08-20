package iirepo_contents_test

import (
	"github.com/reiver/go-iirepo/contents"

	"fmt"
)

func ExamplePath() {

	var parent string = "/home/joeblow/workspaces/myproject"

	repoContentsPath := iirepo_contents.Path(parent)

	fmt.Printf("The repo's contents directory's path is: %s\n", repoContentsPath)

	// Output:
	// The repo's contents directory's path is: /home/joeblow/workspaces/myproject/.ii/contents
}
