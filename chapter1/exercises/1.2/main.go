// Modify the echo program to print the index and
// value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Program name:", os.Args[0])

	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Line", i, "-", os.Args[i])

	}
}
