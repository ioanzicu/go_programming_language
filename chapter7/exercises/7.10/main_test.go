package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tcs := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 1}, false},
		{[]int{1}, true},
		{[]int{1, 1}, true},
		{[]int{1, 3}, false},
	}

	for _, tc := range tcs {
		obtained := IsPalindrome(nums(tc.nums))
		if obtained != tc.expected {
			t.Errorf("%v, expected %v but obtained %v", tc.nums, tc.expected, obtained)
		}
	}
}
