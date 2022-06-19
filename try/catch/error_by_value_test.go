package catch_test

import (
	"fmt"

	"github.com/apitalist/lang/try"
	"github.com/apitalist/lang/try/catch"
)

func ExampleErrorByValue() {
	var myError = fmt.Errorf("my error")

	try.Catch(
		func() {
			// Panic with an error:
			panic(myError)
		},
		// Handle a specific error type:
		catch.ErrorByValue(
			myError,
			func(err error) {
				fmt.Printf("Caught custom error: %v", err)
			},
		),
	)

	// Output: Caught custom error: my error
}
