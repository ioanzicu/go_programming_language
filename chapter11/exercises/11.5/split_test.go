// Extend TestSplit to use a table of inputs and expected outputs.
package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tcs = []struct {
		str     string
		sep     string
		expects int
	}{
		{"x:y:z", ":", 3},
		{"x:y:z", "-", 1},
		{"x-y-z", "-", 3},
		{"x,y,z", ":", 1},
	}
	for _, tc := range tcs {
		words := strings.Split(tc.str, tc.sep)
		if actual := len(words); actual != tc.expects {
			t.Errorf("Split(%q, %q), expects: %d, actual: %d.", tc.str, tc.sep, tc.expects, actual)
		}
	}
}
