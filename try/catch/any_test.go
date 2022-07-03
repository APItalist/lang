package catch_test

import (
	"fmt"

	"github.com/apitalist/lang/v2/try"
	"github.com/apitalist/lang/v2/try/catch"
)

func ExampleAny() {
	try.Catch(
		func() {
			// Panic with an error:
			panic("something bad happened")
		},
		// Handle a specific error type:
		catch.Any(
			func(err any) {
				fmt.Printf("Caught panic: %v", err)
			},
		),
	)

	// Output: Caught panic: something bad happened
}
