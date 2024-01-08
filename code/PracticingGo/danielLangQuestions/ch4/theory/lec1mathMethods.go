package main

import (
	"fmt"
	"math"
)

func angleToRadian(degree float64) float64 {
	return degree / 360 * 2 * math.Pi
}

func radianToAngle(degreeInRadian float64) float64 {
	return 180 / math.Pi * degreeInRadian
}

func main() {
	// Trigonometric Methods
	fmt.Println("Sine of π/3:", math.Sin(math.Pi/3))
	fmt.Println("Cosine of π/3:", math.Cos(math.Pi/3))
	fmt.Println("Tangent of π/4:", math.Tan(math.Pi/4))

	fmt.Println("Arcsine of 1/2 in degrees:", radianToAngle(math.Asin(float64(1)/2)))
	fmt.Println("Arccosine of 1/2 in degrees:", radianToAngle(math.Acos(float64(1)/2)))
	fmt.Println("Arctangent of 1 in degrees:", radianToAngle(math.Atan(1)))

	// Exponent Methods
	fmt.Println("e^1:", math.Exp(1))
	fmt.Println("Natural logarithm of e:", math.Log(math.E))
	fmt.Println("Log base 10 of 100:", math.Log10(100))
	fmt.Println("2^3:", math.Pow(2, 3))
	fmt.Println("Square root of 2:", math.Sqrt(2))

	// Rounding methods
	fmt.Println("Ceiling of 2.1:", math.Ceil(2.1))
	fmt.Println("Floor of 2.1:", math.Floor(2.1))
	fmt.Println("Round to even of 3.5:", math.RoundToEven(3.5))
	fmt.Println("Round of 3.5:", math.Round(3.5))

	// Min, Max, Abs
	fmt.Println("Maximum of 2 and 3:", math.Max(2, 3))
	fmt.Println("Minimum of 2 and 3:", math.Min(2, 3))
	fmt.Println("Absolute value of -2:", math.Abs(-2))

	// Random methods
	fmt.Println()
}
