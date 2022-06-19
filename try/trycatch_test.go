package try_test

import (
	"fmt"

	"github.com/apitalist/lang/try"
	"github.com/apitalist/lang/try/catch"
)

type myCustomErrorType struct {
}

func (m myCustomErrorType) Error() string {
	return "My custom error"
}

var myOtherError = fmt.Errorf("Some custom error")

func Example() {
	try.Catch(
		func() {
			fmt.Println("Hello world!")
		},
		catch.Any(
			func(err any) {
				fmt.Printf("An error happened: %v", err)
			},
		),
	)
	// Output: Hello world!
}

func ExampleCatch() {
	try.Catch(
		func() {
			// Panic with an error:
			var err error = &myCustomErrorType{}
			panic(err)
		},
		// Handle a specific error type:
		catch.ErrorByType(
			func(err *myCustomErrorType) {
				fmt.Printf("Caught my custom error: %v", err)
			},
		),
		// Handle an error by value:
		catch.ErrorByValue(
			myOtherError,
			func(err error) {
				fmt.Printf("Caught error by value: %v", err)
			},
		),
		// Fallback to catch all errors:
		catch.Error(
			func(err error) {
				fmt.Printf("Caught generic error: %v", err)
			},
		),
		// Fallback to catch anything:
		catch.Any(
			func(err any) {
				fmt.Printf("Caught panic: %v", err)
			},
		),
	)

	// Output: Caught my custom error: My custom error
}

func ExampleCatchHandler() {
	try.Catch(
		func() {
			panic(fmt.Errorf("something bad happened"))
		},
		// Specify a custom catch function:
		func(e any) bool {
			// Return true here if the function can handle the panic.
			return false
		},
		// Or use the built-in functions:
		catch.ErrorByType(
			func(err *myCustomErrorType) {
				fmt.Printf("Caught my custom error: %v", err)
			},
		),
		catch.ErrorByValue(
			myOtherError,
			func(err error) {
				fmt.Printf("Caught error by value: %v", err)
			},
		),
		catch.Error(
			func(err error) {
				fmt.Printf("Caught generic error: %v", err)
			},
		),
		catch.Any(
			func(err any) {
				fmt.Printf("Caught panic: %v", err)
			},
		),
	)

	// Output: Caught generic error: something bad happened
}
