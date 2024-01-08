package main

import (
	"fmt"
	"time"
)

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}
}

func solve(currTime time.Time, offset int) time.Time {
	return currTime.Add(time.Duration(int(time.Hour) * offset))
}

func main() {
	value := time.Now().UTC()
	var offset int
	fmt.Print("Enter the time zone offset to GMT:")
	takeInput(&offset)
	value = solve(value, offset)
	// layout is according to predefined RFCs, use them to define which value is needed where
	fmt.Println("The current time is:", value.Format("15💀04💀05😁"))
}
