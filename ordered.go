package lang

// Ordered is an interface that collects all primitives that can be ordered. This is typically helpful when creating
// Go 1.18 generic definitions:
//
//     func compare[T lang.Ordered](a, b T) int {
//     }
//
// In the example above, a and b can be compared with the smaller than (<) and larger than (>) operators despite not
// knowing their types. Ideally, Golang will, at some point, add a native interface for this and we will deprecate this
// extension.
type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}
