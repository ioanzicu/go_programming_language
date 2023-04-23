package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"surface/eval"
	"surface/surface"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}

	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}

	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance form (0, 0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

/*
http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)

http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/12

http://localhost:8000/plot?expr=sin(x*y/10)/10
*/

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
