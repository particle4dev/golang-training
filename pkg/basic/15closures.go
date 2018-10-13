package basic

// URL 	https://gobyexample.com/closures

// IntSeq return an anonymous functions
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
