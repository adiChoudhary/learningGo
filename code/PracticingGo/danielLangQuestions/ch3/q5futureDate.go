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

func weekDay(num int) int {
	return num % 7
}

func main() {
	var inputDate, timeElapsed int
	takeInput(&inputDate)
	takeInput(&timeElapsed)
	fmt.Printf("%v %v", time.Weekday(weekDay(inputDate)), time.Weekday(weekDay(inputDate+timeElapsed)))
}
