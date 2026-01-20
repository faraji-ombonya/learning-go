// Defer

// Defer is used to ensure that a function call is perfomed later in a
// program's execution, usually for purposes of cleanup. defer is
// often used where e.g. ensure and finally would be used in other languages

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Suppose we wanted to create a file, write to itm and then close when
// we're done. Here's how we could do that with defer.

func main() {
	path := filepath.Join(os.TempDir(), "defer.txt")
	f := createFile(path)

	// Immediately after getting a file object with createFile, we defer the
	// closing of tha file with closeFile. This will be executed at the end
	// of the enclosing function (main), after writeFile has finished.
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Println(f, "data")
}

// It's important to check for errors when closing a file, even in a
// deferred function.
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		panic(err)
	}
}

// Running the program confirms that the file is closed after being written
