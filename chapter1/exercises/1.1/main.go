// Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println("Program name:", os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
