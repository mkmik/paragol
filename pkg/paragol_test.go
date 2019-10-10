package paragol_test

import (
	"fmt"

	paragol "github.com/mkmik/paragol/pkg"
)

func ExampleIsBalanced() {
	examples := []string{
		"(1)",
		"(1",
		"(1)(2)",
	}
	for _, e := range examples {
		fmt.Printf("%q -> %v\n", e, paragol.IsBalanced(e))
	}

	// Output:
	// "(1)" -> true
	// "(1" -> false
	// "(1)(2)" -> false
}
