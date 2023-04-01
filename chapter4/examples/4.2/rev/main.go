// reverse reverses a slice of ints in place.
package main

import "fmt"

func main() {
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// reverse(a[:])  // passes arr as a slice [:]
	// fmt.Println(a) // [5 4 3 2 1 0]

	// Rotate to the left by two positions.
	s := []int{0, 1, 2, 3, 4, 5}

	// reverse(s[:2])
	// fmt.Println(s) // "[1 0 2 3 4 5]"

	// reverse(s[2:])
	// fmt.Println(s) // "[1 0 5 4 3 2]"

	// reverse(s)
	// fmt.Println(s) // "[2 3 4 5 0 1]"

	// Rotate to the right
	reverse(s)
	fmt.Println(s) // "[5 4 3 2 1 0]"

	reverse(s[:2])
	fmt.Println(s) // "[4 5 3 2 1 0]"

	reverse(s[2:])
	fmt.Println(s) // "[4 5 0 1 2 3]"

	x := []string{"1", "2", "3"}
	y := []string{"1", "2", "5"}

	z := []string{"1", "2", "3"}
	fmt.Printf("Are equal %v\t| x=%v | y=%v\n", equal(x, y), x, y)
	fmt.Printf("Are equal %v\t| x=%v | z=%v\n", equal(x, z), x, z)

	// fmt.Println("")
	// var b []int    // len(s) == 0, s == nil
	// b = nil        // len(s) == 0, s == nil
	// b = []int(nil) // len(s) == 0, s == nil
	// b = []int{}    // len(s) == 0, s != nil

	fmt.Println("")
	a := make([]int, 5)
	fmt.Printf("a=%v\n", a)

	b := make([]int, 10)[:5]
	fmt.Printf("b=%v\n", b)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// compare 2 string slices
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
