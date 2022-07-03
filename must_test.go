package lang_test

import (
	"fmt"

	"github.com/apitalist/lang/v2"
)

func ExampleMust() {
	funcReturningError := func() error {
		return nil
	}

	// If funcReturningError returns an error, lang.Must will throw a panic.
	lang.Must(funcReturningError())
	fmt.Println("No panic thrown.")

	// Output: No panic thrown.
}

func ExampleMust2() {
	funcReturningError := func() (string, error) {
		return "Hello world!", nil
	}

	// If funcReturningError returns an error, lang.Must2 will throw a panic.
	message := lang.Must2(funcReturningError())
	fmt.Println(message)

	// Output: Hello world!
}

func ExampleMust3() {
	funcReturningError := func() (string, string, error) {
		return "Hello", "world!", nil
	}

	// If funcReturningError returns an error, lang.Must3 will throw a panic.
	message1, message2 := lang.Must3(funcReturningError())
	fmt.Println(message1, message2)

	// Output: Hello world!
}

func ExampleMust4() {
	funcReturningError := func() (string, string, string, error) {
		return "Hello", "world", "!", nil
	}

	// If funcReturningError returns an error, lang.Must4 will throw a panic.
	message1, message2, message3 := lang.Must4(funcReturningError())
	fmt.Println(message1, message2, message3)

	// Output: Hello world !
}
