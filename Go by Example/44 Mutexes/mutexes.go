// Mutexes

// In the previous example we saw how to manage simple counter state using
// atomic operations. For more complex state we can use a mutex to safely
// access data accross multiple goroutines.

package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to updte it concurrently
// from multiple go routines, we add a Mutex to synchronize access. Note that
// mutexes must not be copied, so of struct is passed arounf, it should be
// done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	// Lock the mutex before accessing counters; unlock it at the end of the
	// function using a defer statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {

	c := Container{
		// Note the zero value of a mutex is usable as-is, so no initialization
		// is required here.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// This function increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	// Run several increments a named concurrently; note that the all access
	// the same container, and two of  them access the same counter.
	wg.Go(func() { doIncrement("a", 10000) })
	wg.Go(func() { doIncrement("a", 10000) })
	wg.Go(func() { doIncrement("b", 10000) })

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}

// Running the program shows that the counters updated as expected.

// Next we'll look at implementing this same state management task using only
// goroutines and channels.
