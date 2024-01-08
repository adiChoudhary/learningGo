package main

import (
	"fmt"
	"math/rand"
)

func saySomething(whatToSay func(int) bool) {
	if whatToSay(20) {
		fmt.Println("Its Awesome")
	} else {
		fmt.Println("Its Hell")
	}
}

func boolGenerator(seed int) bool {
	num := rand.Intn(100)
	fmt.Println(num)
	return num+seed > 50
}

func main() {
	// We can easily pass functions as values in go
	saySomething(boolGenerator)
}
