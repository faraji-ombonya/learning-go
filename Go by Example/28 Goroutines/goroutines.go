// Goroutines

// A goroutine is a lightweight thread of execution

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function cvall f(s). Here's how we'd call that in the
	// usual way, running it sychronously.
	f("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine
	// will execute concurrently eith the calling one.
	go f("goroutine")

	// You can also start a goroutinr for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in seperate goroutines
	// now.
	time.Sleep(time.Second)
	fmt.Println("done")
}

// When we run this program, we see the outpu of the blocking call first, then
// the output of the two goroutines. The goroutines' output may be interleaved
// because goroutines are being run concurrently by the Go runtime.
