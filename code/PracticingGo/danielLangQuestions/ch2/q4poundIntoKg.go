package main

import "fmt"

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}
}

func main() {
	var pound float32
	fmt.Print("Enter a number in pounds:")
	takeInput(&pound)
	const PoundToKg = 0.454
	fmt.Printf("%g pounds is %g kg", pound, PoundToKg*pound)
}
