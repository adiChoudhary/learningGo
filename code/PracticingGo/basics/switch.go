package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	// Here we don't require a break statement since go evaluates only the case which evaluates to be true
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	switch num := "11"; num {
	case strconv.Itoa(10):
		fmt.Println("Its A")
	default:
		fmt.Println(num)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("%T\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch without a condition provides clean way to write long if else ladder
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
