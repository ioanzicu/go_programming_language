package main

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tcs := []struct {
		expr     string
		env      Env
		expected string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, tc := range tcs {
		// Print expr only when it changes
		if tc.expr != prevExpr {
			fmt.Printf("\n%s\n", tc.expr)
			prevExpr = tc.expr
		}
		expr, err := Parse(tc.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		actual := fmt.Sprintf("%.6g", expr.Eval(tc.env))
		fmt.Printf("\t%v => %s\n", tc.env, actual)
		if actual != tc.expected {
			t.Errorf("%s.Eval() in %v = %q, expected %q\n", tc.expr, tc.env, actual, tc.expected)
		}
	}
}
