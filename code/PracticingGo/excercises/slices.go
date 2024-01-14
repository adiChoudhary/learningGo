package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	picSlice := make([][]uint8, dy)
	for i := 0; i < len(picSlice); i++ {
		picSlice[i] = make([]uint8, dx)
		for j := 0; j < len(picSlice[i]); j++ {
			if i == j {
				picSlice[i][j] = 128
			}
		}
	}
	return picSlice
}

func main() {
	fmt.Println(Pic(10, 10))
}
