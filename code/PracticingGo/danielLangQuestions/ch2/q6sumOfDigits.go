package main

import (
	"fmt"
	"math/rand"
)

// eg of naked function in Go
func sumOfDigits(num int) (sum int) {
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return
}

func main() {
	num := rand.Intn(1000)
	fmt.Printf("The sum of digits of %d is %d", num, sumOfDigits(num))
}
