/*
Rewrite the PopCount example from Section 2.6.2 so that it initializes the lookup
table using sync.Once the first time it is needed. (Realistically, the cost of synchronization
would be prohibitive for a small and highly optimized function like PopCount.)
*/

package popcount

import "sync"

// pc[i] is the population count of i.

var pc [256]byte
var pcs [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop returns the population count (number of set bits) of x by looping
func PopCountLoop(x uint64) int {
	n := 0
	for i := uint(0); i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

// PopCountTable returns the population count (number of set bits) of x by table-lookup
var initTableOnce sync.Once

func initSync() {
	initTableOnce.Do(func() {
		for i := range pc {
			pcs[i] = pcs[i/2] + byte(i&1)
		}
	})
}

func PopCountTableSync(x uint64) int {
	initSync()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
