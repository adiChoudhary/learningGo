package main

import "fmt"

// Similar to java, go also assign default values to unassigned variables
// These are known as 'Zero Values' in go's terminology
var i, c, java bool

func main() {
	var num int
	fmt.Println(i, c, java, num)

	// variables can also be initialized as given below
	// we can omit type if we are initializing a variable
	var x, y int = 1, 2
	fmt.Println(x, y)

	// inside a function, we can omit the var keyword and use shortHand ':=' for initialization
	// For untyped numeric constants, type is decided by that value
	a, b, c, d := "Wow", 1, true, 3.8
	fmt.Println(a, b, c, d)
}
