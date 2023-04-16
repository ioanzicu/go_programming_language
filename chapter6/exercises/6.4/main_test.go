package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	if x.String() != "{1 9 144}" {
		t.Errorf("expected {1 9 144}, got %s", x.String())
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 144, 9)
	if x.String() != "{1 9 144}" {
		t.Errorf("expected {1 9 144}, got %s", x.String())
	}
	if x.Len() != 3 {
		t.Error("expected length = 3")
	}
}

func TestHas(t *testing.T) {
	var x IntSet
	x.Add(1)
	if !x.Has(1) {
		t.Error("expected to contain 1")
	}
}

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	if x.Len() != 3 {
		t.Error("expected length = 3")
	}
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.UnionWith(&y)

	if x.Len() != 4 {
		t.Error("expected length = 4")
	}

	if x.String() != "{1 9 42 144}" {
		t.Errorf("expected {1 9 42 144}, got %s", x.String())
	}
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(63) // if is > 63 it will fail
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.IntersectWith(&y)

	if x.Len() != 1 {
		t.Error("expected length = 1")
	}

	if x.String() != "{9}" {
		t.Errorf("expected {9}, got %s", x.String())
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(63) // if is > 63 it will fail
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.DifferenceWith(&y)

	if x.Len() != 2 {
		t.Error("expected length = 2")
	}

	if x.String() != "{1 63}" {
		t.Errorf("expected {1 63}, got %s", x.String())
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(63) // if is > 63 it will fail
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.SymmetricDifference(&y)

	if x.Len() != 3 {
		t.Error("expected length = 3")
	}

	if x.String() != "{1 42 63}" {
		t.Errorf("expected {1 42 63}, got %s", x.String())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	if x.String() != "{1 9 144}" {
		t.Errorf("expected {1 9 144}, got %s", x.String())
	}
	if x.Len() != 3 {
		t.Error("expected length = 3")
	}

	x.Remove(1)
	if x.String() != "{9 144}" {
		t.Errorf("expected {9 144}, got %s", x.String())
	}
	if x.Len() != 2 {
		t.Error("expected length = 2")
	}

	x.Remove(9)
	if x.String() != "{144}" {
		t.Errorf("expected {144}, got %s", x.String())
	}
	if x.Len() != 1 {
		t.Error("expected length = 1")
	}

	x.Remove(144)
	if x.String() != "{}" {
		t.Errorf("expected {}, got %s", x.String())
	}
	if x.Len() != 0 {
		t.Error("expected length = 0, got", x.Len())
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	x.Clear()
	if x.String() != "{}" {
		t.Errorf("expected {}, got %s", x.String())
	}
	if x.Has(1) {
		t.Error("expected to not have 1")
	}
	if x.Has(1) {
		t.Error("expected to not have 1")
	}
	if x.Has(9) {
		t.Error("expected to not have 1")
	}

	if x.Len() != 0 {
		t.Error("expected length = 0, got", x.Len())
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y := x.Copy()
	if y.Len() != x.Len() {
		t.Error("expected length = 3, got", y.Len(), x.Len())
	}

	// x.Add(33)

	if y.String() != x.String() {
		t.Errorf("expected %v to be equal to %v", y.String(), x.String())
	}
}

func TestElems(t *testing.T) {
	var x IntSet

	exp := []int{1, 9, 144}
	x.AddAll(exp...)

	for i, item := range x.Elems() {
		if item != exp[i] {
			t.Errorf("%d. expected %v to be equal to %v", i, item, exp[i])
		}
	}
}
