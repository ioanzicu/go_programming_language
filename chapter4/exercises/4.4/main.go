// Write a version of rotate that operates in a single pass.
package main

import "fmt"

func main() {
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// reverse(a[:])  // passes arr as a slice [:]
	// fmt.Println(a) // [5 4 3 2 1 0]

	// Rotate to the left by two positions.
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s) // [0 1 2 3 4 5]
	rotate(s, 3)
	fmt.Println(s) // "[3 4 5 0 1 2]"

	fmt.Println("")

	ss := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(ss) // [0 1 2 3 4 5]
	ss = rotate2(ss, 3)
	fmt.Println(s) // "[3 4 5 0 1 2]"
}

// Time O(n) Space O(n)
// modify the argument slice
func rotate(s []int, i int) {
	temp := make([]int, 0)
	// append the first half
	temp = append(temp, s[i:]...)
	// append the second half
	temp = append(temp, s[:i]...)
	// copy the result back
	copy(s, temp)
}

// returns a new slice
func rotate2(s []int, i int) []int {
	var temp []int
	// append the first half
	temp = append(temp, s[i:]...)
	// append the second half
	temp = append(temp, s[:i]...)
	// copy the result back
	return temp
}
