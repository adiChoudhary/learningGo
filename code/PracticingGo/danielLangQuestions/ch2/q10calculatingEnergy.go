package main

import "fmt"

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func main() {
	const HEAT_CONSTANT = 4184
	var finalTemp, initialTemp, m float32
	takeInput(&m)
	takeInput(&initialTemp)
	takeInput(&finalTemp)
	q := m * (finalTemp - initialTemp) * HEAT_CONSTANT
	fmt.Println(q)
}
