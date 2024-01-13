package main

import "fmt"

func sumTwoNum(x int, y int) int {
	return x + y
}

// Type can be omitted except from last if variables share datatype
func sumTwoNumBetter(x, y int) int {
	return x + y
}

// A go function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Naked return statement(Just a go concept)
// Generally used only for short functions because it impacts the readability of the program
func nakedReturn(sum int) (x, y int) {
	x = sum + 10
	y = sum - 10
	return
}

func main() {
	fmt.Println(sumTwoNum(2, 3))
	fmt.Println(sumTwoNumBetter(2, 3))
	fmt.Println(swap("FirstString", "SecondString"))
	fmt.Println(nakedReturn(12))
}
