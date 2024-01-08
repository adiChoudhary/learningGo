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

func main() {
	var num int
	takeInput(&num)
	randomInt := rand.Intn(2)
	fmt.Println(randomInt)
	zeroCnt := 0
	oneCnt := 0
	for _ = range make([]int, 100000) {
		temp := rand.Intn(2)
		if temp == 0 {
			zeroCnt++
		} else {
			oneCnt++
		}
	}
	fmt.Printf("Zero cnt:%d One cnt:%d\n", zeroCnt, oneCnt)
	if randomInt == num {
		fmt.Println("Yeah winner")
	} else {
		fmt.Println("Sore losser")
	}
}
