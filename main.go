package main

import(
	"fmt"
	"math"
)

func GenDisplaceFn(a float64, v float64, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5 * a * math.Pow(t, 2) + (v*t) + s
	}
}

func main() {
	var a, v, s, t float64

	fmt.Print("Enter acceleration (a): ")
	fmt.Scan(&a)

	fmt.Print("Enter initial velocity (v): ")
	fmt.Scan(&v)

	fmt.Print("Enter initial displacement (s): ")
	fmt.Scan(&s)

	fn := GenDisplaceFn(a, v, s)

	fmt.Print("Enter time (t): ")
	fmt.Scan(&t)

	displacement := fn(t)
	fmt.Printf(" The displacement after %.2f seconds is: %.2f\n", t, displacement)
}