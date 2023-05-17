/*
Write benchmarks for Add, UnionWith, and other methods of *IntSet (§6.5)
using large pseudo-random inputs. How fast can you make these methods run?
How does the choice of word size affect performance? - Time Linearly - O(n) Space O(n) - map, O(1) - IntSet
How fast is IntSet compared to a set implementation based on the built-in map type?
On average 2X, and IntSet has practically no memory allocation

go test -v -bench=. -benchmem
=== RUN   TestLenZeroInitially
--- PASS: TestLenZeroInitially (0.00s)
=== RUN   TestLenAfterAddingElements
--- PASS: TestLenAfterAddingElements (0.00s)
=== RUN   TestRemove
--- PASS: TestRemove (0.00s)
=== RUN   TestClear
--- PASS: TestClear (0.00s)
=== RUN   TestCopy
--- PASS: TestCopy (0.00s)
=== RUN   TestAddAll
--- PASS: TestAddAll (0.00s)
goos: linux
goarch: amd64
pkg: intset
cpu: Intel(R) Core(TM) i5-10310U CPU @ 1.70GHz
BenchmarkMapIntSetAdd10
BenchmarkMapIntSetAdd10-8                1896138               553.6 ns/op           323 B/op          3 allocs/op
BenchmarkMapIntSetAdd100
BenchmarkMapIntSetAdd100-8                163322              6785 ns/op            3490 B/op         19 allocs/op
BenchmarkMapIntSetAdd1000
BenchmarkMapIntSetAdd1000-8                15384             76663 ns/op           53445 B/op         74 allocs/op
BenchmarkMapIntSetHas10
BenchmarkMapIntSetHas10-8               42970921                25.02 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetHas100
BenchmarkMapIntSetHas100-8              41060209                28.31 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetHas1000
BenchmarkMapIntSetHas1000-8             42700210                29.49 ns/op            0 B/op          0 allocs/op
BenchmarkMapIntSetAddAll10
BenchmarkMapIntSetAddAll10-8             2535805               457.7 ns/op           323 B/op          3 allocs/op
BenchmarkMapIntSetAddAll100
BenchmarkMapIntSetAddAll100-8             190893              5292 ns/op            3493 B/op         19 allocs/op
BenchmarkMapIntSetAddAll1000
BenchmarkMapIntSetAddAll1000-8             17062             70566 ns/op           53451 B/op         74 allocs/op
BenchmarkMapIntSetString10
BenchmarkMapIntSetString10-8              801704              1335 ns/op             360 B/op         15 allocs/op
BenchmarkMapIntSetString100
BenchmarkMapIntSetString100-8              78325             16190 ns/op            4561 B/op        108 allocs/op
BenchmarkMapIntSetString1000
BenchmarkMapIntSetString1000-8              6331            170307 ns/op           41124 B/op        991 allocs/op
BenchmarkBitIntSetAdd10
BenchmarkBitIntSetAdd10-8                4807814               242.1 ns/op             0 B/op          0 allocs/op
BenchmarkBitIntSetAdd100
BenchmarkBitIntSetAdd100-8                531556              2297 ns/op               0 B/op          0 allocs/op
BenchmarkBitIntSetAdd1000
BenchmarkBitIntSetAdd1000-8                59223             21041 ns/op               0 B/op          0 allocs/op
BenchmarkBitIntSetHas10
BenchmarkBitIntSetHas10-8               56420022                18.55 ns/op            0 B/op          0 allocs/op
BenchmarkBitIntSetHas100
BenchmarkBitIntSetHas100-8              59200196                18.70 ns/op            0 B/op          0 allocs/op
BenchmarkBitIntSetHas1000
BenchmarkBitIntSetHas1000-8             62024625                18.52 ns/op            0 B/op          0 allocs/op
BenchmarkBitIntSetAddAll10
BenchmarkBitIntSetAddAll10-8            21010945                57.65 ns/op            0 B/op          0 allocs/op
BenchmarkBitIntSetAddAll100
BenchmarkBitIntSetAddAll100-8            4450984               263.8 ns/op             0 B/op          0 allocs/op
BenchmarkBitIntSetAddAll1000
BenchmarkBitIntSetAddAll1000-8            494794              2461 ns/op               0 B/op          0 allocs/op
BenchmarkBitIntSetString10
BenchmarkBitIntSetString10-8              813928              1402 ns/op             256 B/op         13 allocs/op
BenchmarkBitIntSetString100
BenchmarkBitIntSetString100-8              93202             13368 ns/op            3641 B/op        106 allocs/op
BenchmarkBitIntSetString1000
BenchmarkBitIntSetString1000-8             10000            114519 ns/op           32914 B/op        991 allocs/op
PASS
ok      intset  31.818s
*/
package main

import (
	"math/rand"
	"testing"
)

func newIntSets() []IntSet {
	return []IntSet{&BitIntSet{}, NewMapIntSet()}
}

func TestLenZeroInitially(t *testing.T) {
	for _, s := range newIntSets() {
		if s.Len() != 0 {
			t.Errorf("%T.Len(): got %d, want 0", s, s.Len())
		}
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(2000)
		if s.Len() != 2 {
			t.Errorf("%T.Len(): got %d, want 2", s, s.Len())
		}
	}
}

func TestRemove(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Remove(0)
		if s.Has(0) {
			t.Errorf("%T: want zero removed, got %s", s, s)
		}
	}
}

func TestClear(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(1000)
		s.Clear()
		if s.Has(0) || s.Has(1000) {
			t.Errorf("%T: want empty set, got %s", s, s)
		}
	}
}

func TestCopy(t *testing.T) {
	for _, orig := range newIntSets() {
		orig.Add(1)
		copy := orig.Copy()
		copy.Add(2)
		if !copy.Has(1) || orig.Has(2) {
			t.Errorf("%T: want %s, got %s", orig, orig, copy)
		}
	}
}

func TestAddAll(t *testing.T) {
	for _, s := range newIntSets() {
		s.AddAll(0, 2, 4)
		if !s.Has(0) || !s.Has(2) || !s.Has(4) {
			t.Errorf("%T: want {2 4}, got %s", s, s)
		}
	}
}

const max = 32000

func addRandom(set IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(max))
	}
}

func benchHas(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.Has(rand.Intn(max))
	}
}

func benchAdd(b *testing.B, set IntSet, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
		set.Clear()
	}
}

func randInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(max)
	}
	return ints
}

func benchAddAll(b *testing.B, set IntSet, batchSize int) {
	ints := randInts(batchSize)
	for i := 0; i < b.N; i++ {
		set.AddAll(ints...)
		set.Clear()
	}
}

func benchUnionWith(bm *testing.B, a, b IntSet, n int) {
	addRandom(a, n)
	addRandom(b, n)
	for i := 0; i < bm.N; i++ {
		a.UnionWith(b)
	}
}

func benchString(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.String()
	}
}

//	func Benchmark<Type><Method><Size>(b *testing.B) {
//		bench<Method>(b, New<Type>(), <Size>)
//	}
func BenchmarkMapIntSetAdd10(b *testing.B)      { benchAdd(b, NewMapIntSet(), 10) }
func BenchmarkMapIntSetAdd100(b *testing.B)     { benchAdd(b, NewMapIntSet(), 100) }
func BenchmarkMapIntSetAdd1000(b *testing.B)    { benchAdd(b, NewMapIntSet(), 1000) }
func BenchmarkMapIntSetHas10(b *testing.B)      { benchHas(b, NewMapIntSet(), 10) }
func BenchmarkMapIntSetHas100(b *testing.B)     { benchHas(b, NewMapIntSet(), 100) }
func BenchmarkMapIntSetHas1000(b *testing.B)    { benchHas(b, NewMapIntSet(), 1000) }
func BenchmarkMapIntSetAddAll10(b *testing.B)   { benchAddAll(b, NewMapIntSet(), 10) }
func BenchmarkMapIntSetAddAll100(b *testing.B)  { benchAddAll(b, NewMapIntSet(), 100) }
func BenchmarkMapIntSetAddAll1000(b *testing.B) { benchAddAll(b, NewMapIntSet(), 1000) }
func BenchmarkMapIntSetString10(b *testing.B)   { benchString(b, NewMapIntSet(), 10) }
func BenchmarkMapIntSetString100(b *testing.B)  { benchString(b, NewMapIntSet(), 100) }
func BenchmarkMapIntSetString1000(b *testing.B) { benchString(b, NewMapIntSet(), 1000) }
func BenchmarkBitIntSetAdd10(b *testing.B)      { benchAdd(b, NewBitIntSet(), 10) }
func BenchmarkBitIntSetAdd100(b *testing.B)     { benchAdd(b, NewBitIntSet(), 100) }
func BenchmarkBitIntSetAdd1000(b *testing.B)    { benchAdd(b, NewBitIntSet(), 1000) }
func BenchmarkBitIntSetHas10(b *testing.B)      { benchHas(b, NewBitIntSet(), 10) }
func BenchmarkBitIntSetHas100(b *testing.B)     { benchHas(b, NewBitIntSet(), 100) }
func BenchmarkBitIntSetHas1000(b *testing.B)    { benchHas(b, NewBitIntSet(), 1000) }
func BenchmarkBitIntSetAddAll10(b *testing.B)   { benchAddAll(b, NewBitIntSet(), 10) }
func BenchmarkBitIntSetAddAll100(b *testing.B)  { benchAddAll(b, NewBitIntSet(), 100) }
func BenchmarkBitIntSetAddAll1000(b *testing.B) { benchAddAll(b, NewBitIntSet(), 1000) }
func BenchmarkBitIntSetString10(b *testing.B)   { benchString(b, NewBitIntSet(), 10) }
func BenchmarkBitIntSetString100(b *testing.B)  { benchString(b, NewBitIntSet(), 100) }
func BenchmarkBitIntSetString1000(b *testing.B) { benchString(b, NewBitIntSet(), 1000) }

func BenchMarkMapIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 10)
}

func BenchMarkMapIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 100)
}
func BenchMarkMapIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 1000)
}
func BenchMarkBitIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 10)
}
func BenchMarkBitIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 100)
}
func BenchMarkBitIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 1000)
}
