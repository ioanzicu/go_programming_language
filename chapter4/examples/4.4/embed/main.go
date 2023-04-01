package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point  Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 8
	// w.Circle.Radius = 5
	// w.Spokes = 33

	// w = Wheel{Circle{Point: 8, 8}, 5}, 33}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 33,
	}

	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:33}

	w.Circle.Point.X = 42 // when imported should work w.X

	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:42, Y:8}, Radius:5}, Spokes:33}
}
