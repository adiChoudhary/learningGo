package main

import (
	"fmt"
	"strconv"
)

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}
}

func numToDollarString(num float32) string {
	return "$" + strconv.FormatFloat(float64(num), 'g', -1, 32)
}

func main() {
	var (
		subtotal float32
		gratuity float32
	)
	fmt.Print("Enter the subtotal and a gratuity rate:")
	takeInput(&subtotal)
	takeInput(&gratuity)
	gratuityAmount := subtotal * gratuity / 100
	total := subtotal + gratuityAmount
	fmt.Printf("The gratuity is %s and total is %s", numToDollarString(gratuityAmount), numToDollarString(total))
}
