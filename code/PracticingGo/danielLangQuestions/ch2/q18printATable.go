package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i <= 5; i++ {
		if i == 0 {
			fmt.Printf("%-10s%-10s%-10s\n", "a", "b", "pow(a, b)")
		} else {
			fmt.Printf("%-10d%-10d%-10d\n", i, i+1, int(math.Pow(float64(i), float64(i+1))))
		}
	}
}
