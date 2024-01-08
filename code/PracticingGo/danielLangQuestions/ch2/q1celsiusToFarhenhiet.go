package main

import "fmt"

func main() {
	var num float32
	fmt.Print("Enter temperature in Celsius:")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Printf("Error occured %v\n", err)
	}
	fahrenhiet := (9.0/5)*num + 32
	fmt.Printf("%g celsius is %g fahrenhiet", num, fahrenhiet)
}
