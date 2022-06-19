package catch

import (
    "errors"

    "github.com/apitalist/lang/try"
)

// ErrorByType creates an error handler that catches custom type errors. For example, errors are typically declared
// as a custom struct like this:
//
//     type myCustomErrorType struct {
//     }
//
//     func (m *myCustomErrorType) Error() string {
//         return "This is a custom error"
//     }
func ErrorByType[E error](f func(err E)) try.CatchHandler {
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
