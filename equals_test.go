package lang_test

import "fmt"

type complexObject struct {
	number int
}

func (c *complexObject) Equals(other *complexObject) bool {
	return c.number == other.number
}

func ExampleEquals() {
	var a, b *complexObject
	a = &complexObject{
		1,
	}
	b = &complexObject{
		2,
	}

	if a.Equals(b) {
		fmt.Println("A equals B")
	} else {
		fmt.Println("A does not equal B")
	}

	// Output: A does not equal B
}
