package lang

import "errors"

// TryCatch runs the specified function (f). If the function throws a panic, the handlers are tried in order to handle
// the panic. The handler functions can be custom, or one of the built-in error handlers, such as ErrorType, ErrorValue,
// ErrorAny, or AnyPanic
func TryCatch(
	f func(),
	catch1 Handler,
	catch ...Handler,
) {
	result := run(f)
	if result != nil {
		for _, h := range append([]Handler{catch1}, catch...) {
			if h(result) {
				return
			}
		}
	}
	panic(result)
}

func run(f func()) (result interface{}) {
	defer func() {
		result = recover()
	}()
	f()
	return
}

// Handler is a function that can process a panic value.
type Handler func(e any) bool

// ErrorValue creates an error handler that catches fixed value errors. For example, errors are typically declared
// outside any functions like this:
//
//     var ErrMyError = fmt.Errorf("some value")
func ErrorValue(value error, f func(err error)) Handler {
	return func(e any) bool {
		err, ok := e.(error)
		if !ok {
			return false
		}
		if errors.Is(err, value) {
			f(err)
			return true
		}
		return false
	}
}

// ErrorType creates an error handler that catches custom type errors. For example, errors are typically declared
// as a custom struct like this:
//
//     type myCustomErrorType struct {
//     }
//
//     func (m *myCustomErrorType) Error() string {
//         return "This is a custom error"
//     }
func ErrorType[E error](f func(err E)) Handler {
	return func(e any) bool {
		err, ok := e.(error)
		if !ok {
			return false
		}
		var errorType E
		if errors.As(err, &errorType) {
			f(errorType)
			return true
		}
		return false
	}
}

// ErrorAny creates a handler that catches any error types. This is typically used as a last error handler.
func ErrorAny(f func(err error)) Handler {
	return func(e any) bool {
		err, ok := e.(error)
		if !ok {
			return false
		}
		f(err)
		return true
	}
}

// AnyPanic creates a catch-all error handler.
func AnyPanic(f func(err any)) Handler {
	return func(e any) bool {
		f(e)
		return true
	}
}
