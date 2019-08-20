package iirepo_test

import (
	"github.com/reiver/go-iirepo"

	"fmt"
)

func ExampleName() {

	repoName := iirepo.Name()

	fmt.Printf("The repo directory's name is: %s\n", repoName)

	// Output:
	// The repo directory's name is: .ii
}
