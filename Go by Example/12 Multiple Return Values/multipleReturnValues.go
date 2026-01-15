// Multiple Return Values

// Go has built-in support for multiple return values. This feature is used
// often in idiomatic Go, for example to return both the result and errror
// values from a function.

package main

import "fmt"

// The (int, int) in this function signature shows that the function returns
// 2 ints
func vals()(int, int) {
	return 3, 7
}

func main() {
	// Here we use the two different return values from the call with multiple
	// assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you want a subset of the returned value use the blank identifier _.
	_, c := vals()
	fmt.Println(c)
}