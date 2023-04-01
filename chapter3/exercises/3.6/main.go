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
		xmin         = -2
		ymin         = -2
		xmax         = +2
		ymax         = +2
		width        = 1024
		height       = 1024
		doubleWidth  = width * 2
		doubleHeight = height * 2
	)

	f, err := os.Create("out.gif")
	defer f.Close()
	if err != nil {
		fmt.Printf("unable to create file, %v", err)
		os.Exit(1)
	}

	var colors [doubleWidth][doubleHeight]color.Color

	for py := 0; py < doubleHeight; py++ {
		y := float64(py)/doubleHeight*(ymax-ymin) + ymin
		for px := 0; px < doubleWidth; px++ {
			x := float64(px)/doubleWidth*(xmax-xmin) + xmin
			z := complex(x, y)
			colors[px][py] = mandelbrot(z)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// supersampling
			px1, py1 := 2*px, 2*py

			r1, g1, b1, a1 := colors[px1][py1].RGBA()
			r2, g2, b2, a2 := colors[px1+1][py1].RGBA()
			r3, g3, b3, a3 := colors[px1+1][py1+1].RGBA()
			r4, g4, b4, a4 := colors[px1][py1+1].RGBA()

			avgColor := color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 4),
				uint8((g1 + g2 + g3 + g4) / 4),
				uint8((b1 + b2 + b3 + b4) / 4),
				uint8((a1 + a2 + a3 + a4) / 4),
			}

			img.Set(px, py, avgColor)
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
