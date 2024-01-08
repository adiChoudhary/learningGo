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

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func mpInitializeManyToOne(mp map[int]int, keys []int, value int) {
	for _, reqKey := range keys {
		mp[reqKey] = value
	}
}

func main() {
	/*
		Number of Days in a month
		31 Days: 1 3 5 7 8 10 12
		28/29 Days: 2 (Find if year is leap year or not)
		30 Days: 2 4 6 9
	*/
	var month, year int
	takeInput(&month)
	takeInput(&year)
	mp := make(map[int]int)
	mpInitializeManyToOne(mp, []int{1, 3, 5, 7, 8, 10, 12}, 31)
	mpInitializeManyToOne(mp, []int{4, 6, 9, 11}, 30)
	if isLeapYear(year) {
		mpInitializeManyToOne(mp, []int{2}, 29)
	} else {
		mpInitializeManyToOne(mp, []int{2}, 28)
	}
	fmt.Printf("%s has %d days", time.Month(month), mp[month])
}
