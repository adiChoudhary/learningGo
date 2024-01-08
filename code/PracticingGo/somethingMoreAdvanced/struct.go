package main

import "fmt"

type Vertex struct {
	x int
	y int
}

type AnotherVertex struct {
	X, Y int
}

func main() {
	data := Vertex{1, 2}
	fmt.Println(data)
	data.x = 10
	fmt.Println(data)

	// using pointers to struct
	pointerData := &data
	anotherPtr := &data
	// both type of annotations are valid here
	fmt.Printf("%v %v\n", pointerData.x, (*anotherPtr).y)

	// All given below initializations are valid
	var (
		v1  = AnotherVertex{1, 2}
		v2  = AnotherVertex{X: 1}
		v3  = AnotherVertex{}
		ptr = &Vertex{3, 4}
	)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(*ptr)
}
