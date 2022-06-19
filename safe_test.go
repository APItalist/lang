package lang_test

import (
	"errors"
	"fmt"

	"github.com/apitalist/lang"
)

func ExampleSafe() {
	err := lang.Safe(
		func() {
			panic("oh no")
		},
	)
	if err != nil {
		var t lang.SafeNotErrorPanic
		if errors.As(err, &t) {
			fmt.Printf("An error happened: %v", t.Cause())
		}
	}

	// Output: An error happened: oh no
}
