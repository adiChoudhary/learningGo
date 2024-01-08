package main

import "fmt"

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func main() {
	var num int
	takeInput(&num)
	revNum := 0
	for temp := num; temp > 0; temp /= 10 {
		revNum *= 10
		revNum += temp % 10
	}
	if revNum == num {
		fmt.Printf("%d is a palindrome", num)
	} else {
		fmt.Printf("%d is a not palindrome", num)
	}
}
