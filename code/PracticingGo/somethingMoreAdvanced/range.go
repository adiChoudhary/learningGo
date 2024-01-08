package main

import "fmt"

func main() {
	//pow := []int{1, 2, 4, 8, 16, 32, 64}
	//for i, v := range pow {
	//	fmt.Printf("%d %d\n", i, v)
	//}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << i // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
