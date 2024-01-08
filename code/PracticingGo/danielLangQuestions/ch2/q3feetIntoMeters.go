package main

import "fmt"

// here data has the pointer of the value to be read
func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}
}

func main() {
	var feet float32
	fmt.Print("Enter a value for feet:")
	takeInput(&feet)
	fmt.Printf("%g feet is %g meter", feet, 0.305*feet)
}
