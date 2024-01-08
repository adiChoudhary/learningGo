package main

import "fmt"

func main() {
	a := 7
	pointerA := &a
	// type of pointer to int -> *int
	fmt.Printf("a = %v, &a = %v, typeOf &a = %T\n", a, pointerA, pointerA)
	*pointerA = *pointerA * 10
	fmt.Printf("a = %d\n", a)
}
