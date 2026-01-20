// Timers

// We often want to execute Go code at some point in the future, or repeatedly
// at some interval. Go's built-it timer and ticker features make both of these
// tasks easy. We'l look first at timers and then tickers.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers represent a single event in the future. You tell the timer howlong
	// you want to wait, and it provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer's channel C until it sends a value
	// indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used time.Sleep. One reason a
	// timer may be useful is that you can cancel the timer before it fires.
	// Here's an example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give the timer2 enough time to fire, if it was ever going to,to show it
	// is in fact stopped.
	time.Sleep(2 * time.Second)
}

// The first timer will fire ~2s after we start the program, but the second
// should be stopped before it has a chance to fire.
