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
	//var amount, rate float32
	var inputTimeStr string
	takeInput(&inputTimeStr)

	// https://yourbasic.org/golang/format-parse-string-time-date-example/#layout-options
	inputTime, err := time.Parse("02/01/2006", inputTimeStr)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
	fmt.Println(inputTime)
}
