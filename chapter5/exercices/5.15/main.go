// Write variadic functions max and min, analogous to sum.
// What should these functions do when called with no arguments?
// Write variants that require at least one argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	// min() // error: "please provide at least 1 argument"
	// max() // error: "please provide at least 1 argument"

	fmt.Println(min(1))          // 1
	fmt.Println(min(1, 2))       // 1
	fmt.Println(min(1, 2, 3))    // 1
	fmt.Println(min(333, 3, 33)) // 3

	fmt.Println()

	fmt.Println(max(1))          // 1
	fmt.Println(max(1, 2))       // 2
	fmt.Println(max(1, 2, 3))    // 3
	fmt.Println(max(333, 3, 33)) // 333
}

func min(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("please provide at least 1 argument")
		os.Exit(1)
	}
	minVal := vals[0]
	for _, val := range vals {
		if minVal > val {
			minVal = val
		}
	}
	return minVal
}

func max(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("please provide at least 1 argument")
		os.Exit(1)
	}

	maxVal := vals[0]
	for _, val := range vals {
		if maxVal < val {
			maxVal = val
		}
	}
	return maxVal
}
