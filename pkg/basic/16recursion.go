package basic

// URL 	https://gobyexample.com/closures

// ArithmeticSequence return the sum of the first n terms of an arithmetic sequence
func ArithmeticSequence(end, term int) int {
	if end <= 0 {
		return 0
	}
	return end + ArithmeticSequence(end-term, term)
}
