package basic

// URL 	https://gobyexample.com/functions
// 		https://gobyexample.com/multiple-return-values
// 		https://gobyexample.com/variadic-functions

// Plus takes two ints and returns their sum as an int.
func Plus(a int, b int) int {
	return a + b
}

// PlusPlus takes three ints and returns their sum as an int.
func PlusPlus(a, b, c int) int {
	return a + b + c
}

// Vals multiple return values
func Vals() (int, int) {
	return 3, 7
}

// Sum takes many ints and returns their sum as an int.
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
