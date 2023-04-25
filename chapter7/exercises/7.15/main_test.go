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

func TestString(t *testing.T) {
	tcs := []struct {
		expr   string
		env    Env
		result string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
	}
	var prevExpr string
	for _, tc := range tcs {
		// Print expr only when it changes.
		if tc.expr != prevExpr {
			fmt.Printf("\n%s\n", tc.expr)
			prevExpr = tc.expr
		}
		expr, err := Parse(tc.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		s := expr.String()
		// parsed again,
		fmt.Println("Expr!!!!!!:", s)
		reexpr, err := Parse(s)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// yield an equivalent tree.
		got := fmt.Sprintf("%.6g", expr.Eval(tc.env))
		regot := fmt.Sprintf("%.6g", reexpr.Eval(tc.env))

		//		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != tc.result || regot != tc.result {
			t.Errorf("\n%s.Eval() in %v = %q,\n%s.Eval() in %v = %q, result %q\n",
				tc.expr, tc.env, got, reexpr, tc.env, regot, tc.result)
		}
	}
}

func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		env   Env
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"x % 2", nil, "unexpected '%'"},
		{"!true", nil, "unexpected '!'"},
		{"log(10)", nil, `unknown function "log"`},
		{"sqrt(1, 2)", nil, "call to sqrt has 2 args, want 1"},
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"4!", nil, "24"},
	}

	for _, test := range tests {
		expr, err := Parse(test.input)
		if err == nil {
			err = expr.Check(map[Var]bool{})
		}
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%s: got %q, want %q", test.input, err, test.want)
			}
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s: %v => %s, want %s",
				test.input, test.env, got, test.want)
		}
	}
}
