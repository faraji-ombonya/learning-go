// Atomic Counters

// The primary mechanism for managing state in Go is communication over
// channels. We saw this for example with worker pools. There are a few
// other options for managing state thoug. Here we'll lool at using the
// sync/atomic package for atomic counters accessed by multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// We'll use an atomic integer type to represent our (always-positive)
	// counter.
	var ops atomic.Uint64

	// A waitGroup wi;; help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// We'll start 50 goroutines that each increment the counter exactly
	// 1000 times.
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				// To atomically increment the counter we use Add.
				ops.Add(1)
			}
		})
	}

	// Wait untill all the goroutines are done.
	wg.Wait()

	// Here no goroutines are writing to 'ops', but using Load it's safe to
	// atomically read a value even while other goroutines are (atomically)
	// updating it.
	fmt.Println("ops:", ops.Load())
}

// We xpect to get exactly 50,000 operations. Had we used a non atomic integer
// and incermented it with ops++, we'd likely get a different number, changing
// between runsm because the goroutines would interfere with each other.
// Moreover, we'd get data race failures when running with the -race flag.

// Next we'll look at mutexes, another tool for managing state.
