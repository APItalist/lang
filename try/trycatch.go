package try

// Catch runs the specified function (f). If the function throws a panic, the handlers are tried in order to handle
// the panic. The handler functions can be custom, or one of the built-in error handlers, such as CatchErrorByType,
// CatchErrorByValue, CatchAnyError, or CatchAny.
func Catch(
	f func(),
	catch1 CatchHandler,
	catch ...CatchHandler,
) {
	result := run(f)
	if result != nil {
		for _, h := range append([]CatchHandler{catch1}, catch...) {
			if h(result) {
				return
			}
		}
		panic(result)
	}
}

func run(f func()) (result interface{}) {
	defer func() {
		result = recover()
	}()
	f()
	return
}

// CatchHandler is a function that can process a panic value. It is used in conjunction with AndCatch.
type CatchHandler func(e any) bool
