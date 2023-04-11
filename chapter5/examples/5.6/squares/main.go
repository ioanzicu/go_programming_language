package main

import "fmt"

// squares returns a function that returns next square number each time it is called.

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
