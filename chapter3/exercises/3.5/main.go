// Implement a full-color Mandelbrot set using the function image.NewRGBA and
// the type color.RGBA or color.YCbCr.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin   = -2
		ymin   = -2
		xmax   = +2
		ymax   = +2
		width  = 1024
		height = 1024
	)

	f, err := os.Create("out.gif")
	defer f.Close()
	if err != nil {
		fmt.Printf("unable to create file, %v", err)
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	err = png.Encode(f, img)
	if err != nil {
		fmt.Printf("unable to encode image to Stdout, %v", err)
		os.Exit(1)
	}
}

// var palette = []color.Color{color.White, color.Black, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}
// RGB
var palette = []color.Color{color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}

// RGB mandelbrot
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = uint8(15)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n%uint8(len(palette))]
		}
	}

	return color.Black
}
