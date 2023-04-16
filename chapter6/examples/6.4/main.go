package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.X-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset)
		path[i] = op(path[i], offset)
	}
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance   // method value
	fmt.Println(distanceFromP(q)) // "5"

	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979"

	scaleP := p.ScaleBy // mtehod value
	scaleP(2)           // p becomes (2, 4)
	scaleP(3)           //      then (6, 12)
	scaleP(10)          //      then (60, 120)

	// method expression
	distance := Point.Distance   // receiver T
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy // receiver *T
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
}
