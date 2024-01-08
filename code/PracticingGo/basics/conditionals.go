package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	var num int
	num = 10
	//scanln, err := fmt.Scanln(&num)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	// scanln has value 1 if valid input else 0
	// err shows the correct value
	//fmt.Println(scanln)

	// Parentheses is necessary
	if num == 10 {
		fmt.Println("😁 You're lucky")
	} else {
		fmt.Println("😭 You're unlucky")
	}

	// Just like for loop, you can put a short assignment expression before a conditional
	// this v is available both inside if and else blocks
	if v := num + 10; v > 30 {
		fmt.Println("You're goat")
	} else {
		fmt.Println("You're a sore loser")
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
