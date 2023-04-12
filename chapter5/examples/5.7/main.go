package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())              // 0
	fmt.Println(sum(3))             // 3
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15

	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...)) // 15

	// The type of function is different than an ordinary slice parameter
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"

	linenum := 12
	// name := []int{2}
	name := "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
}

// variadic function ...
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
