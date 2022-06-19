package lang_test

import (
	"fmt"

	"github.com/apitalist/lang"
)

type myCustomErrorType struct {
}

func (m myCustomErrorType) Error() string {
	return "My custom error"
}

var myOtherError = fmt.Errorf("Some custom error")

func ExampleTryCatch() {
	lang.TryCatch(
		func() {
			// Panic with an error:
			var err error = &myCustomErrorType{}
			panic(err)
		},
		// Handle a specific error type:
		lang.ErrorType(
			func(err *myCustomErrorType) {
				fmt.Printf("Caught my custom error: %v", err)
			},
		),
		// Handle an error by value:
		lang.ErrorValue(
			myOtherError,
			func(err error) {
				fmt.Printf("Caught error by value: %v", err)
			},
		),
		// Fallback to catch all errors:
		lang.ErrorAny(
			func(err error) {
				fmt.Printf("Caught generic error: %v", err)
			},
		),
		// Fallback to catch anything:
		lang.AnyPanic(
			func(err any) {
				fmt.Printf("Caught panic: %v", err)
			},
		),
	)

	// Output: Caught my custom error: My custom error
}

func ExampleHandler() {
	lang.TryCatch(
		func() {
			panic(fmt.Errorf("something bad happened"))
		},
		// Specify a custom handler:
		func(e any) bool {
			// Return true here if the function can handle the panic.
			return false
		},
		// Or use the built-in handlers:
		lang.ErrorType(
			func(err *myCustomErrorType) {
				fmt.Printf("Caught my custom error: %v", err)
			},
		),
		lang.ErrorValue(
			myOtherError,
			func(err error) {
				fmt.Printf("Caught error by value: %v", err)
			},
		),
		lang.ErrorAny(
			func(err error) {
				fmt.Printf("Caught generic error: %v", err)
			},
		),
		lang.AnyPanic(
			func(err any) {
				fmt.Printf("Caught panic: %v", err)
			},
		),
	)

	// Output: Caught generic error: something bad happened
}

func ExampleErrorType() {
	lang.TryCatch(
		func() {
			// Panic with an error:
			var err error = &myCustomErrorType{}
			panic(err)
		},
		// Handle a specific error type:
		lang.ErrorType(
			func(err *myCustomErrorType) {
				fmt.Printf("Caught my custom error: %v", err)
			},
		),
	)

	// Output: Caught my custom error: My custom error
}

func ExampleErrorValue() {
	var myError = fmt.Errorf("my error")

	lang.TryCatch(
		func() {
			// Panic with an error:
			panic(myError)
		},
		// Handle a specific error type:
		lang.ErrorValue(
			myError,
			func(err error) {
				fmt.Printf("Caught custom error: %v", err)
			},
		),
	)

	// Output: Caught custom error: my error
}

func ExampleErrorAny() {
	lang.TryCatch(
		func() {
			// Panic with an error:
			panic(fmt.Errorf("something bad happened"))
		},
		// Handle a specific error type:
		lang.ErrorAny(
			func(err error) {
				fmt.Printf("Caught error: %v", err)
			},
		),
	)

	// Output: Caught error: something bad happened
}

func ExampleAnyPanic() {
	lang.TryCatch(
		func() {
			// Panic with an error:
			panic("something bad happened")
		},
		// Handle a specific error type:
		lang.AnyPanic(
			func(err any) {
				fmt.Printf("Caught panic: %v", err)
			},
		),
	)

	// Output: Caught panic: something bad happened
}
