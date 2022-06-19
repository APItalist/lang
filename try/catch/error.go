package catch

import (
    "github.com/apitalist/lang"
)

// Error creates a handler that catches any error types. This is typically used as a last error handler.
func Error(f func(err error)) lang.Catch {
    return func(e any) bool {
        err, ok := e.(error)
        if !ok {
            return false
        }
        f(err)
        return true
    }
}
