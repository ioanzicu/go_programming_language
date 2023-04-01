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

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(num)
	}
}

/*
go test -bench=.

goos: linux
goarch: amd64
pkg: popcount
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkPopCount-8             1000000000               0.2942 ns/op
BenchmarkPopCountLoop-8         281996464                4.304 ns/op
PASS
ok      popcount        1.980s

*/
