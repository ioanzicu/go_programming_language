// Change the Lissajous program’s color palette to green on black, for added
// authenticity. To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xff},
// where each pair of hexadecimal digits represents the intensity of the red, green, or blue com-
// ponent of the pixel.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{255, 0, 0, 50}, color.RGBA{0, 255, 0, 0}, color.RGBA{0, 0, 255, 0}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	redIndex   = 2
	greenIndex = 3
	blueIndex  = 4
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		// number of complete x oscillator revolutions
		res  = 0.001 // angular resolution
		size = 100
		// image canvas covers [-size..+size]
		nframes = 64
		// number of animation frames
		delay = 8
	// delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				redIndex) // change colors to get different animations
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
