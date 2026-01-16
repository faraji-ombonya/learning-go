// Recursion

package main

import "fmt"

// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	// Anonymous function can also be recursive, but this requires expliciltly
	// declaring a variable with var to store the function before it is
	// defined.
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Since fib was previously declared in main, Go knows which function
		// to call with fib here
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
