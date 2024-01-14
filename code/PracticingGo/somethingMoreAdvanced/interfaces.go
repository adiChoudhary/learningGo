package main

import "fmt"

type Person struct {
	name string
	age  string
}

func (receiver *Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %s\n", receiver.name, receiver.age)
}

func main() {
	a := &Person{"Aditya", "26"}
	fmt.Println(a)
}
