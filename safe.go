package lang

import (
	"fmt"

	"github.com/apitalist/lang/try"
)

// Safe runs the specified function and, if a panic occurs with an error, returns the error. If a panic happens with
// a non-error type, the panic is wrapped in a SafeNotErrorPanic
func Safe(f func()) (err error) {
	try.Catch(
		f,
		func(e any) bool {
			var ok bool
			err, ok = e.(error)
			if !ok {
				err = SafeNotErrorPanic{
					e,
				}
			}
			return true
		},
	)
	return
}

// SafeNotErrorPanic is an error type that is returned if a Safe function catches a non-error panic.
type SafeNotErrorPanic struct {
	cause any
}

// Cause returns the originally caught panic.
func (s SafeNotErrorPanic) Cause() any {
	return s.cause
}

// Error returns the text with the panic.
func (s SafeNotErrorPanic) Error() string {
	return fmt.Sprintf("non-error panic: %v", s.cause)
}
