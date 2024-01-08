package main

import (
	"fmt"
	"time"
)

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func main() {
	var inputMonth int
	takeInput(&inputMonth)
	fmt.Println(time.Month(inputMonth))
}
