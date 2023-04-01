// Write a function that counts the number of bits that are different in two SHA256 hashes
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	count := 0
	for i, x := range c1 {
		x2 := c2[i]
		if PopCount(uint64(x)) != PopCount(uint64(x2)) {
			count++
		}
		// fmt.Printf("%+#v | c1 = %+#v |  %+#v | c2 = %+#v |  %+#v\n", i, x, PopCount(uint64(x)), x2, PopCount(uint64(x2)))
	}
	fmt.Println("Number of bits that are different in c1 and c2 is = ", count)
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
