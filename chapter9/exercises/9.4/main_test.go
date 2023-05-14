package main

import "testing"

func benchmarkPipeline(b *testing.B, stages int) {
	in, out := pipeline(stages)
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

func BenchmarkPipeline10(b *testing.B)       { benchmarkPipeline(b, 10) }
func BenchmarkPipeline100(b *testing.B)      { benchmarkPipeline(b, 100) }
func BenchmarkPipeline1000(b *testing.B)     { benchmarkPipeline(b, 1000) }
func BenchmarkPipeline10000(b *testing.B)    { benchmarkPipeline(b, 10000) }
func BenchmarkPipeline100000(b *testing.B)   { benchmarkPipeline(b, 100000) }
func BenchmarkPipeline1000000(b *testing.B)  { benchmarkPipeline(b, 1000000) }
func BenchmarkPipeline10000000(b *testing.B) { benchmarkPipeline(b, 10000000) } // out of memory

/* 32 GB RAM
go test -v -bench=.

goos: linux
goarch: amd64
pkg: pipeline
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkPipeline10
BenchmarkPipeline10-8             592422              1774 ns/op
BenchmarkPipeline100
BenchmarkPipeline100-8             63295             17400 ns/op
BenchmarkPipeline1000
BenchmarkPipeline1000-8             5398            189387 ns/op
BenchmarkPipeline10000
BenchmarkPipeline10000-8             403           3094711 ns/op
BenchmarkPipeline100000
BenchmarkPipeline100000-8             32          32695434 ns/op
BenchmarkPipeline1000000
BenchmarkPipeline1000000-8             1        1640942227 ns/op
BenchmarkPipeline10000000
signal: killed
FAIL    pipeline        40.628s




With RESET TIMER

goos: linux
goarch: amd64
pkg: pipeline
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkPipeline10
BenchmarkPipeline10-8             602413              1936 ns/op
BenchmarkPipeline100
BenchmarkPipeline100-8             58099             17677 ns/op
BenchmarkPipeline1000
BenchmarkPipeline1000-8             5704            192186 ns/op
BenchmarkPipeline10000
BenchmarkPipeline10000-8             405           3017147 ns/op
BenchmarkPipeline100000
BenchmarkPipeline100000-8             42          30659519 ns/op
BenchmarkPipeline1000000
BenchmarkPipeline1000000-8             4         313340565 ns/op
BenchmarkPipeline10000000
signal: killed
FAIL    pipeline        51.834s

*/
