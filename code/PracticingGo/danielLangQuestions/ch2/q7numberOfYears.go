package main

import "fmt"

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}
}

func minutesToYears(minutes int) (int, int) {
	totalHours := minutes / 60
	totalDays := totalHours / 24
	days := totalDays % 365
	years := totalDays / 365
	return years, days
}

func main() {
	var minutes int
	fmt.Print("Enter the number of minutes:")
	takeInput(&minutes)
	years, days := minutesToYears(minutes)
	fmt.Printf("%d minutes is approximately %d years and %d days\n", minutes, years, days)
}
