package main

import "fmt"

func greatest() func(int) int {
	cur := -1
	// This is an anonymous function with access to variable declared outside its scope also known as closure
	return func(x int) int {
		if x > cur {
			cur = x
			return x
		} else {
			return cur
		}
	}
}

// Uses of closures: https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/

func main() {
	ints := [5]int{1, 2, 3, 4, 5}
	findGreatest := greatest()
	for _, v := range ints {
		fmt.Println(findGreatest(v))
	}
}
