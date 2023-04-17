package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tcs := []struct {
		input    []int
		expected string
	}{
		{
			[]int{2, 3, 1},
			"[1 2 3]",
		},
		{
			[]int{},
			"[0]",
		},
		{
			[]int{1},
			"[1]",
		},
		{
			[]int{22, 33, 100},
			"[22 33 100]",
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			root := new(tree)
			for i, v := range tc.input {
				if i == 0 {
					root = &tree{value: v}
				} else {
					root = add(root, v)
				}
			}
			actual := root.String()
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", actual, root.String())
			}
		})
	}
}
