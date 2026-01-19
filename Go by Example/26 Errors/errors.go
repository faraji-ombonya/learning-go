// Errors

// In Go, it's idiomatic to communicate errors via an explicit separate return
// value. This contrasts with the exceptions used in languages like Java,
// Python and Ruby and the overloaded single result / error value sometimes
// used in C. Go's approach makes it easy to see which functions return errors
// and to handle them using the same language constructs employed for other non
// error tasks.

package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have a type error, a
// built-in interface.
func f(arg int) (int, error) {
	if arg == 42 {

		// errors.New constructs a basic error value with the give error
		// message
		return -1, errors.New("Can't work with 42")
	}

	// A nil value in ther error position indicates that there was no error.
	return arg + 3, nil
}

// A sentinel error is a predeclared variable that is ued to signify a specific
// error condition.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		// We can wrap errors with higher-level errors to add context. The
		// simplest way to do this is with the %w verb in fmt.Errorf.
		// Wrapped errors create a logical chain (A wraps B, which wraps C,
		// etc.) that can be queried with functions like errors.Is and
		// errors.As.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		// it's idiomatic to use an inline error check in the if line.
		if r, e := f(i); e != nil {
			fmt.Println("f failed", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Println("Unknown error: %s\n", err)
			}

			continue
		}

		fmt.Println("Tea is ready")
	}
}
