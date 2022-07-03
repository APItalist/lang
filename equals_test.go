package lang_test

import (
	"fmt"

	"github.com/apitalist/lang/v2"
)

type complexObject struct {
	number int
}

func (c *complexObject) Equals(other any) bool {
	switch o := other.(type) {
	case *complexObject:
		return c.number == o.number
	case int, int32, int64:
		return c.number == o
	}
	return false
}

func ExampleEquals() {
	var a lang.Equals = &complexObject{
		1,
	}
	b := int64(2)

	if a.Equals(b) {
		fmt.Println("A equals B")
	} else {
		fmt.Println("A does not equal B")
	}

	// Output: A does not equal B
}
