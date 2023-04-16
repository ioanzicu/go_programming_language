package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

// same thing, but as a method of the Point type
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.X-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

/*
The embedded field instructs the compiler to generate additional
wrapper methods that delegate to the declared methods.
*/

// func (p ColoredPoint) Distance(q Point) float64 {
// 	return p.Point.Distance(q)
// }

// func (p *ColoredPoint) ScaleBy(factor float64) {
// 	p.Point.ScaleBy(factor)
// }

type ColoredPPoint struct {
	*Point
	Color color.RGBA
}

func main() {
	// fields
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	// methods
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"

	pp := ColoredPPoint{&Point{1, 1}, red}
	qp := ColoredPPoint{&Point{5, 4}, blue}
	fmt.Println(pp.Distance(*qp.Point)) // "5"
	qp.Point = pp.Point                 //pp and qp now share the same Point
	pp.ScaleBy(2)
	fmt.Println(*pp.Point, *qp.Point) // "{2 2} {2 2}"
}
