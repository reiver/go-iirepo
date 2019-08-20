package iirepo_contents_test

import (
	"github.com/reiver/go-iirepo/contents"

	"fmt"
)

func ExampleName() {

	repoContentsName := iirepo_contents.Name()

	fmt.Printf("The repo contents directory's name is: %s\n", repoContentsName)

	// Output:
	// The repo contents directory's name is: contents
}
