package main

import "fmt"

func main() {
	a := 12
	var b float64 = float64(a)
	// This will give error cause there is no concept of implicit type conversion in go
	//var c float64 = a
	fmt.Println(a, b, c)
}
