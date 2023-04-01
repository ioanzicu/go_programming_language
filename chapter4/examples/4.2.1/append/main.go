// reverse reverses a slice of ints in place.
package main

import "fmt"

func main() {
	// var runes []rune
	// for _, r := range "Hello, |-|" {
	// 	runes = append(runes, r)
	// }
	// fmt.Printf("%q\n", runes)
	// fmt.Printf("%q\n", []rune("Hello, |-|"))

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d   cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	fmt.Println("")

	var b []int
	fmt.Println(b) // []

	b = appendIntElipse(b, 1)
	fmt.Println(b) // [1]

	b = appendIntElipse(b, 2, 3)
	fmt.Println(b) // [1, 2, 3]

	b = appendIntElipse(b, 4, 5, 6)
	fmt.Println(b) // [1, 2, 3, 4, 5, 6]

	b = appendIntElipse(b, b...) // append the slice x
	fmt.Println(b)               // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1 // +1 for the new element
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function
	}
	z[len(x)] = y
	return z
}

func appendIntElipse(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	// ...expand z to al least zlen...
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function
	}
	copy(z[len(x):], y)
	return z
}
