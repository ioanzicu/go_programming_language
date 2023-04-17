// Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4)
// that reveals the sequence of values in the tree.
package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	// nums := []int{4, 2, 1, 5, 2, 7}

	// fmt.Printf("Before sort %v\n", nums)
	// Sort(nums)
	// fmt.Printf("After sort %v\n", nums)
	root := new(tree)
	for i, v := range []int{3, 1, 2} {
		if i == 0 {
			root = &tree{value: v}
		} else {
			root = add(root, v)
		}
	}
	actual := root.String()
	fmt.Println(actual) // "[1 2 3]"
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to the values
// in order and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	order := make([]int, 0)
	order = appendValues(order, t)
	if len(order) == 0 {
		return "[]"
	}

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v", order)
	return b.String()
}
