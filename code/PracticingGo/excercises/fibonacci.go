package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		// third variable to handle f0, f1 state is not required its handled by language
		f1, f0 = f1+f0, f1
		return f0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
