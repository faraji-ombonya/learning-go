// Sorting

// Go's slices package implements sorting for builtins and user defined types.
// We'll look at sorting for builtins first.

package main

import (
	"fmt"
	"slices"
)

func main() {
	// Sorting functions are generic, and work for any ordered built-in type.
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting ints.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints: ", ints)

	// We can also use the slices package to check if a slice is already in
	// sorted order.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
