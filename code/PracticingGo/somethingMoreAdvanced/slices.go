package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	primes := [6]int{2, 3, 5, 7, 9, 13}
	var somePrimes []int = primes[1:4]
	fmt.Printf("Type Of Slice: %T, Value Of Slice: %v\n", somePrimes, somePrimes)
	fmt.Printf("Type Of Array: %T, Value Of Array: %v\n", primes, primes)
	// Length and Capacity of array remains same
	fmt.Println(len(primes))
	fmt.Println(cap(primes))

	// slices are like reference to array, changes done in a slice element will effect parent array
	names := [3]string{"Hello", "World", "Wow"}
	someNames := names[0:2]
	fmt.Println(names, someNames)
	someNames[1] = "Nikal yahan se"
	fmt.Println(names, someNames)

	// slice literal example
	// this also creates an underlying array but doesn't gives us the access to it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// Limits to slice can be omitted(Just like Python)
	testingSliceLimits := []int{2, 3, 5, 7, 11, 13}

	testingSliceLimits = testingSliceLimits[1:4]
	fmt.Println(testingSliceLimits)

	testingSliceLimits = testingSliceLimits[:2]
	fmt.Println(testingSliceLimits)

	testingSliceLimits = testingSliceLimits[1:]
	fmt.Println(testingSliceLimits)

	// Slices have both length and capacity
	// length -> actual length of the slice
	// capacity -> length of underlying array starting from first element in slice
	someSliceYouAre := []int{2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(someSliceYouAre)
	someSliceYouAre = someSliceYouAre[:1]
	printSlice(someSliceYouAre)
	someSliceYouAre = someSliceYouAre[4:6]
	printSlice(someSliceYouAre)
	// Below lines gives error because we try to increase capacity of a slice beyond its limits
	//someSliceYouAre = someSliceYouAre[:5]
	//printSlice(someSliceYouAre)

	// Default slice is nil
	var nilSlice []int
	printSlice(nilSlice)
	if nilSlice == nil {
		fmt.Println("Its a nil slice!!!")
	}

	// Creating dynamic arrays using slice
	dynamicSlice := make([]int, 5)
	printSlice(dynamicSlice)

	dynamicSliceV2 := make([]int, 0, 5) // (type, length, cap)
	printSlice(dynamicSliceV2)
	dynamicSliceV2 = dynamicSlice[1:]
	printSlice(dynamicSliceV2)

	// slices can contain any type including other slices
	// if we specify size we can create 2D array in similar fashion
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Append to a slice
	var incSlice []int
	// It works on nil slices also
	incSlice = append(incSlice, 1)
	printSlice(incSlice)
	// it grows, parent array behind it gets replaced with a new one if its capacity is smaller
	incSlice = append(incSlice, 2)
	printSlice(incSlice)
	// We can add multiple elements at a time
	incSlice = append(incSlice, 3, 4, 5, 6, 7, 8)
	printSlice(incSlice)
	// We can also unpack a slice to append its elements to another slice
	incSlice = append(incSlice, []int{10, 20, 30}...)
	someNewSlice := make([]int, 6)
	incSlice = append(incSlice, someNewSlice...)
	printSlice(incSlice)
}
