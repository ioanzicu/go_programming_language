/*
Write benchmarks to compare the PopCount implementation in Section 2.6.2
with your solutions to Exercise 2.4 and Exercise 2.5.
At what point does the table-based approach break even?

At 10 and 100

 go test -bench=.
goos: linux
goarch: amd64
pkg: popcount
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkPopCountTable1-8                  31766            129017 ns/op
BenchmarkPopCountTable10-8                 10000            352430 ns/op
BenchmarkPopCountTable100-8                10000           3607260 ns/op
BenchmarkPopCountTable1000-8                3216          11985395 ns/op
BenchmarkPopCountTable10000-8                330          11984847 ns/op
BenchmarkPopCountTable100000-8               100          35770544 ns/op
BenchmarkPopCountShift1-8               34474782                34.34 ns/op
BenchmarkPopCountShift10-8               3607024               331.7 ns/op
BenchmarkPopCountShift100-8               360355              3370 ns/op
BenchmarkPopCountShift1000-8               35193             32629 ns/op
BenchmarkPopCountShift10000-8               3463            325884 ns/op
BenchmarkPopCountShift100000-8               357           3276071 ns/op
BenchmarkPopCountClears1-8              366964461                3.155 ns/op
BenchmarkPopCountClears10-8             38875513                31.69 ns/op
BenchmarkPopCountClears100-8             2860537               392.1 ns/op
BenchmarkPopCountClears1000-8             272104              4402 ns/op
BenchmarkPopCountClears10000-8             24908             47724 ns/op
BenchmarkPopCountClears100000-8             2181            546320 ns/op
PASS
ok      popcount        107.203s

*/

package popcount

import "testing"

func benchmarkPopCount(b *testing.B, f func(uint64) int, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func benchmarkPopCountTable(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for j := range pc {
			pc[j] = pc[j/2] + byte(j&1)
		}
		benchmarkPopCount(b, PopCountTable, n)
	}
}

func BenchmarkPopCountTable1(b *testing.B)      { benchmarkPopCountTable(b, 1) }
func BenchmarkPopCountTable10(b *testing.B)     { benchmarkPopCountTable(b, 10) }
func BenchmarkPopCountTable100(b *testing.B)    { benchmarkPopCountTable(b, 100) }
func BenchmarkPopCountTable1000(b *testing.B)   { benchmarkPopCountTable(b, 1000) }
func BenchmarkPopCountTable10000(b *testing.B)  { benchmarkPopCountTable(b, 10000) }
func BenchmarkPopCountTable100000(b *testing.B) { benchmarkPopCountTable(b, 100000) }

func BenchmarkPopCountShift1(b *testing.B)      { benchmarkPopCount(b, PopCountShift, 1) }
func BenchmarkPopCountShift10(b *testing.B)     { benchmarkPopCount(b, PopCountShift, 10) }
func BenchmarkPopCountShift100(b *testing.B)    { benchmarkPopCount(b, PopCountShift, 100) }
func BenchmarkPopCountShift1000(b *testing.B)   { benchmarkPopCount(b, PopCountShift, 1000) }
func BenchmarkPopCountShift10000(b *testing.B)  { benchmarkPopCount(b, PopCountShift, 10000) }
func BenchmarkPopCountShift100000(b *testing.B) { benchmarkPopCount(b, PopCountShift, 100000) }

func BenchmarkPopCountClears1(b *testing.B)      { benchmarkPopCount(b, PopCountClears, 1) }
func BenchmarkPopCountClears10(b *testing.B)     { benchmarkPopCount(b, PopCountClears, 10) }
func BenchmarkPopCountClears100(b *testing.B)    { benchmarkPopCount(b, PopCountClears, 100) }
func BenchmarkPopCountClears1000(b *testing.B)   { benchmarkPopCount(b, PopCountClears, 1000) }
func BenchmarkPopCountClears10000(b *testing.B)  { benchmarkPopCount(b, PopCountClears, 10000) }
func BenchmarkPopCountClears100000(b *testing.B) { benchmarkPopCount(b, PopCountClears, 100000) }
