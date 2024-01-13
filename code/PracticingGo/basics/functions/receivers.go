package main

import "fmt"

type Link struct {
	Text string
	URL  string
}

// Note: Receiver names should be same

// Value receiver
func (receiver Link) printLink() {
	fmt.Println(receiver.URL, receiver.Text)
}

// Pointer receiver
func (receiver *Link) initializeLink() {
	receiver.Text = "Hello World"
	receiver.URL = "helloWorld.com"
}

func main() {
	link := Link{}
	link.initializeLink()
	link.printLink()
}
