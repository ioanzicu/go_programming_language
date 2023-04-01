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

func PopCountClearRightMostBit(x uint64) int {
	n := 0
	for x != 0 {
		n++
		x = x & (x - 1)
	}

	return n
}

/*
go test -bench=.


goos: linux
goarch: amd64
pkg: popcount
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkPopCount-8                     1000000000               0.2797 ns/op
BenchmarkPopCountClearRightMostBit-8    264211080                4.497 ns/op
PASS
ok      popcount        1.966s

*/
