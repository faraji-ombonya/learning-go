package main

import "fmt"

func main() {
	i := 2
	fmt.Print("Write", i, "as ")

	switch i {
	case 1:
		fmt.Println("one")
	}
}
