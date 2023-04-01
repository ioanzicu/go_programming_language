// Write a web server that renders fractals and writes the image data to the client.
// Allow the client to specify the x, y, and zoom values as parameters to the HTTP request

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/mandelbrot", mandelbrotHanlder) // http://localhost:8080/mandelbrot
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func mandelbrotHanlder(w http.ResponseWriter, r *http.Request) {
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	if widthStr == "" || heightStr == "" {
		fmt.Fprintf(w, "Please provide width and height in the query parameters\nExample: mandelbrot?width=1024&height=1024")
		return
	}

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("width", width)
	fmt.Println("height", height)

	if width == 0 {
		width = 1024
	}

	if height == 0 {
		height = 1024
	}
	mandelbrot := getMandelbrotImage(width, height)
	w.Header().Set("Content-Type", "image/gif")
	err = png.Encode(w, mandelbrot)
	if err != nil {
		fmt.Printf("unable to encode image to Stdout, %v", err)
		return
	}
}

func getMandelbrotImage(width, height int) *image.RGBA {

	const (
		xmin = -2
		ymin = -2
		xmax = +2
		ymax = +2
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// image point (px, py) represents complex value z
			img.Set(px, py, mandelbrot(z))
		}
	}

	return img
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
