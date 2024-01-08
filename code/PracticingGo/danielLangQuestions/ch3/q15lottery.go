package main

import (
	"fmt"
	"math/rand"
)

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func numToList(num int) [10]int {
	var mp [10]int
	for num > 0 {
		mp[num%10]++
		num /= 10
	}
	return mp
}

func compAllDigits(num1 int, num2 int) bool {
	// freqMap of each num
	mp1 := numToList(num1)
	mp2 := numToList(num2)
	return mp1 == mp2
}

func compAnyDigit(num1 int, num2 int) bool {
	// freqMap of each num
	mp1 := numToList(num1)
	mp2 := numToList(num2)
	for i := 0; i < len(mp1); i++ {
		if mp1[i] > 0 && mp2[i] > 0 {
			return true
		}
	}
	return false
}

func main() {
	var num int
	takeInput(&num)
	randNum := rand.Intn(900) + 100
	fmt.Printf("Input:%d Random Int:%d\n", num, randNum)
	if num == randNum {
		fmt.Println("Grand Masti")
	} else if compAllDigits(num, randNum) {
		fmt.Println("Masti")
	} else if compAnyDigit(num, randNum) {
		fmt.Println("Its Ok")
	} else {
		fmt.Println("Sore Loser")
	}
}
