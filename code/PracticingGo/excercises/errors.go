package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	//fmt.Println(x.Error())
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// This method is invoked implicitly by formatting methods like
// println if a type has implemented Error method
func (p ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can't Square Root %v", float64(p))
}

func main() {
	fmt.Println(math.Sqrt(20))
	fmt.Println(Sqrt(-20))
}
