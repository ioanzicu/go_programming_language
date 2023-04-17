// Add a method Elems that returns a slice containing the elements
// of the set, suitable for iterating over with a range loop.
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64) // quotient and remainder
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// varadic AddAll
// allows a list of values to be added
func (s *IntSet) AddAll(items ...int) {
	for _, item := range items {
		s.Add(item)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// 10010101
			// |
			// 10001001
			// --------
			// 10011101
			s.words[i] |= tword // OR
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// 10010101
			// &
			// 10001001
			// --------
			// 10000001 // only 1 in both

			s.words[i] &= tword // AND
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// 10010101
			//^01101010
			// &
			// 10001001
			// --------
			// 00001000 // only 1 in both

			s.words[i] &= ^tword // AND XOR
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference sets s to the symmetirc difference of s and t.
// symmetric difference of two sets contains the elements present
// in one set or the other but not both
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			// 10010101
			// ^
			// 10001001
			// --------
			// 00011100 // exclude 1 in both

			s.words[i] ^= tword // XOR
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func bitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

// return the number of elements
func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		l += bitCount(word)
	}
	return l
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		fmt.Printf("cannot remove element %d, because the set does not have such element", x)
	}
	word, bit := x/64, uint(x%64)
	if len(s.words) == 0 || word > len(s.words) {
		return
	}
	// 1 << bit         ex: 00010000
	// ^(1<<bit)			11101111

	// s.words[word]    ex: 10110010
	//  					AND
	// ^(1<<bit)			11101111
	//                      10100010
	s.words[word] &= ^(1 << bit)
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = make([]uint64, len(s.words))
	copy(t.words, s.words)
	return &t
}

// returns a slice containing the elements of the set,
// suitable for iterating over with a range loop.
func (s *IntSet) Elems() []int {
	var items = make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				items = append(items, 64*i+j)
			}
		}
	}
	return items
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())      // "{1, 9, 144}"
	fmt.Println("Len:", x.Len()) // 3

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())      // "{9 42}"
	fmt.Println("Len:", y.Len()) // 2

	x.UnionWith(&y)
	fmt.Println(x.String())      // "{1 9 42 144}"
	fmt.Println("Len:", x.Len()) // 4

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x)         // "{1 9 42 144}"                   *IntSet has method string
	fmt.Println(x.String()) // "{1 9 42 144}"                   compiles implicit inserts &x.String()
	fmt.Println(x)          // "{[4398046511618 0 65536]}"      IntSet does NOT have string method

	for i, item := range x.Elems() {
		fmt.Printf("%d. %d\n", i+1, item)
	}
}
