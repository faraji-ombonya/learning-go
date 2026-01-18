// Channel Synchronization

// We can use channels to synchronize execution across goroutines. Here's an
// example of using a blocking receive to wait for a go routine to finiah. When
// waiting for multiple goroutines to finish, you may prefer to use a
// WaitGroup.

package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The done channel will be used
// to notify another goroutinr that this function's work is done.
func worker(done chan bool) {
	fmt.Print("Wprking...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done
	done <- true
}

func main() {
	done := make(chan bool)

	// Start a worker go routine, giving it the channel to notify on.
	go worker(done)

	// Block untill we receive a notification from the worket on the
	// channel.
	<-done
}

// If you remove the <- done line from this program, the program could
// exit before the worker finished its work, or in some cases even before
// it started.