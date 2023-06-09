// Rewrite PopCount to use a loop instead of a single expression. Compare the per-
// formance of the two versions. (Section 11.4 shows how to compare the performance of differ-
// ent implementations systematically.)
package popcount

// pc[i] is the population count of i.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

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

// PopCount returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	var res byte

	for i := 0; i < 8; i++ {
		res += pc[byte(x>>(i*8))]
	}

	return int(res)
}
