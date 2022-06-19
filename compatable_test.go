package lang_test

import "fmt"

type comparableComplexObject struct {
	data int
}

func (o comparableComplexObject) Compare(b *comparableComplexObject) int {
	return o.data - b.data
}

func ExampleComparable() {
	var a, b *comparableComplexObject
	a = &comparableComplexObject{
		1,
	}
	b = &comparableComplexObject{
		2,
	}

	comparison := a.Compare(b)
	if comparison < 0 {
		fmt.Println("A is smaller than B")
	} else if comparison > 0 {
		fmt.Println("B is smaller than A")
	} else {
		fmt.Println("A and B are equal")
	}

	// Output: A is smaller than B
}
