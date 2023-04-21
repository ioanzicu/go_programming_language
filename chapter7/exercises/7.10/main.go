// The sort.Interface type can be adapted to other uses. Write a function
// IsPalindrome(s sort.Interface) bool that reports whether the sequence s is a palin-
// drome, in other words, reversing the sequence would not change it. Assume that the elements
// at indices i and j are equal if !s.Less(i, j) && !s.Less(j, i).
package main

import (
	"fmt"
	"sort"
)

type nums []int

func (s nums) Len() int {
	return len(s)
}

func (s nums) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s nums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {
	n := nums([]int{1, 2, 3, 2, 1})
	fmt.Println(n, "is palindrome?", IsPalindrome(n))

	n = nums([]int{1, 2, 3, 4, 1})
	fmt.Println(n, "is palindrome?", IsPalindrome(n))
}
