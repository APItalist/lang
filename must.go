package lang

// Must provides error-to-panic conversion on return values.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Must2 provides error-to-panic conversion with one return value.
func Must2[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Must3 provides error-to-panic conversion with two return values.
func Must3[T any, K any](value1 T, value2 K, err error) (T, K) {
	if err != nil {
		panic(err)
	}
	return value1, value2
}

// Must4 provides error-to-panic conversion with three return values.
func Must4[T any, K any, L any](value1 T, value2 K, value3 L, err error) (T, K, L) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3
}
