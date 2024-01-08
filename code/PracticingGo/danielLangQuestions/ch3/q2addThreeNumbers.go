package main

import (
	"fmt"
	"math/rand"
)

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func main() {
	a, b, c := rand.Intn(10), rand.Intn(10), rand.Intn(10)
	//fmt.Println(a + b + c)
	var inputSum int
	tries := 1
	for {
		takeInput(&inputSum)
		actualSum := a + b + c
		switch {
		case actualSum < inputSum:
			fmt.Println("Lower")
		case actualSum > inputSum:
			fmt.Println("Higher")
		case actualSum == inputSum:
			fmt.Printf("You won in %d tries", tries)
			return
		default:
			fmt.Println("Invalid input")
		}
		tries++
	}
}
