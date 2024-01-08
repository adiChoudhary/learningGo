package main

import "fmt"

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func main() {
	var v0, v1, t float32
	fmt.Print("Enter v0, v1, and t:")
	takeInput(&v0)
	takeInput(&v1)
	takeInput(&t)
	a := (v1 - v0) / t
	fmt.Printf("The average acceleration is %g", a)
}
