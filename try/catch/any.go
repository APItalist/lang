package catch

import "github.com/apitalist/lang"

// Any creates a catch-all error handler.
func Any(f func(err any)) lang.Catch {
    return func(e any) bool {
        f(e)
        return true
    }
}
