package lang_test

import (
    "fmt"

    "github.com/apitalist/lang/v1"
)

func compare[T lang.Ordered](a, b T) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } else {
        return 0
    }
}

func ExampleOrdered() {
    a := 1
    b := 2
    if compare(a, b) < 0 {
        fmt.Println("A is smaller than B")
    } else if compare(a, b) > 0 {
        fmt.Println("B is smaller than A")
    } else {
        fmt.Println("A and B are equal")
    }

    // Output: A is smaller than B
}
