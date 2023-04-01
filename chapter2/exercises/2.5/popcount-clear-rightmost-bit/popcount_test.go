// Rewrite PopCount to use a loop instead of a single expression. Compare the per-
// formance of the two versions. (Section 11.4 shows how to compare the performance of differ-
// ent implementations systematically.)
package popcount

import (
	"testing"
)

const num = 100000

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(num)
	}
}

func BenchmarkPopCountClearRightMostBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearRightMostBit(num)
	}
}

/*
go test -bench=.

goos: linux
goarch: amd64
pkg: popcount-shift
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkPopCount-8             1000000000               0.2941 ns/op
BenchmarkPopCountShift-8        63639615                19.55 ns/op
PASS
ok      popcount-shift  1.594s

*/
