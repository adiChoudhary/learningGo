package main

import "fmt"

func main() {
	// Similar to lambda's in other languages, go has anonymous functions which can help write inline code

	// Directly calling an anonymous functions
	func() {
		fmt.Println("Hello World")
	}()
	// Can also be used in this way
	temp := func(x int) bool {
		return x > 30
	}
	fmt.Println(temp(2))
	fmt.Println(temp(41))
}
