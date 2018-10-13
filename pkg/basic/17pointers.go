package basic

// URL 	https://gobyexample.com/pointers

// We’ll show how pointers work in contrast to values with 2 functions:
// NumberVal and NumberPoi.
// NumberVal has an int parameter, so arguments will be passed to it by value.
// NumberVal will get a copy of ival distinct from the one in the calling function.
// NumberPoi in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the value at the referenced address.

// NumberVal set n to 1
func NumberVal(ival int) {
	ival = 1
}

// NumberPoi set n to 10
func NumberPoi(iptr *int) {
	*iptr = 10
}
