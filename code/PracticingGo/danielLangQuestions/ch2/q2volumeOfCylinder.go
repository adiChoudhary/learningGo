package main

import (
	"fmt"
	"math"
)

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}
}

func main() {
	var (
		radius float32
		length float32
	)
	// for passing values by reference use pointer
	fmt.Print("Enter the radius and length of a cylinder:")
	takeInput(&radius)
	takeInput(&length)
	area := radius * radius * math.Pi
	volume := area * length
	// Here .6 tells total digits that will be displayed
	fmt.Printf("The area is %.6g\n", area)
	fmt.Printf("The volume is %.6g", volume)
}
