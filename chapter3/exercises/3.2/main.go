// Experiment with visualizations of other functions from the math package. Can
// you produce an egg box, moguls, or a saddle?
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange...+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	f, err := os.Create("out.svg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	f.WriteString(s)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isFinite(ax) || isFinite(ay) ||
				isFinite(bx) || isFinite(by) ||
				isFinite(cx) || isFinite(cy) ||
				isFinite(dx) || isFinite(dy) {
				continue
			}
			s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			f.WriteString(s)
		}
	}
	f.WriteString("</svg>")
	f.Close()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := moguls(x, y)
	// moguls
	// eggBox
	// snowPile

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,xy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func snowPile(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 10
}

func moguls(x, y float64) float64 {
	return math.Pow(math.Sin(x/xyrange*3*math.Pi), 2) * math.Cos(y/xyrange*3*math.Pi)
}

func saddle(x, y float64) float64 {
	return math.Pow(y/xyrange*2, 2) - math.Pow(x/xyrange*2, 2)
}

func isFinite(f float64) bool {
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return true
	}
	return false
}
