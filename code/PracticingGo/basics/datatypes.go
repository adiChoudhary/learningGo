package main

import (
	"fmt"
	"math/cmplx"
)

// below method to declare variables in go is known as factoring
var (
	num         int        = 2
	unsignedNum uint64     = 1<<64 - 1
	isTrue      bool       = true
	name        string     = "Aditya"
	pi          float32    = 3.14
	complexNum  complex128 = cmplx.Sqrt(-5 + 12i)
)

// Something to learn, add content once done

// Format Specifiers
// %T -> to check datatype of a variable
// %v -> to check value of a variable
func printFormatted(data interface{}) {
	fmt.Printf("Type: %T, Value:%v\n", data, data)
}

func main() {
	printFormatted(num)
	printFormatted(unsignedNum)
	printFormatted(isTrue)
	printFormatted(name)
	printFormatted(pi)
	printFormatted(complexNum)
}
