// Constants

// Go supports constant of character, string, boolean, and numeric values.

package main

import (
	"fmt"
	"math"
)

// const declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A const statement can also appear inside a function body.
	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type untill it's given one, such as by an
	// explicit conversion
	fmt.Println(int64(d))

	// A number can be given a type by using it in context that requires one,
	// su as a variable assignment or function call. For example, here
	// math.Sin expects a float64
	fmt.Println(math.Sin(n))
}
