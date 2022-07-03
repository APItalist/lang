package lang

// Equals defines the Equals method to compare the current object to another object. It does not, however, give you the
// ability to compare two objects with the equals (=) operator, which will always compare by value.
type Equals interface {
	// Equals compares the current object with the passed other object and returns true if the two objects are logically
	// equal.
	Equals(other any) bool
}
