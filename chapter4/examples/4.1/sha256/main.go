// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	a := [1]int{-1}
	arr(a)
	fmt.Printf("last %+#v\n", a)

}

func arr(a [1]int) { // is copied
	fmt.Printf("bef %+#v\n", a)
	a[0] = 100
	fmt.Printf("aft %+#v\n", a)
}
