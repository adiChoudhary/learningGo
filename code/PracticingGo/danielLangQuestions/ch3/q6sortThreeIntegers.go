package main

import "fmt"

func takeInput(data interface{}) {
	_, err := fmt.Scan(data)
	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	}
}

func main() {
	var a, b, c int
	takeInput(&a)
	takeInput(&b)
	takeInput(&c)
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	fmt.Printf("%d %d %d", a, b, c)
}
