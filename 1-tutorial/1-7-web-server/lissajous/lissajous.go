// Generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

// that's a composite literal - in this case a slice
var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 1}}

// const values are fixed at compile time
// scoping also affects these variables
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer) {
	const (
		cycles  = 10    // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 1000  // image canvas covers [-size..+size]
		nframes = 100   // number of animation frames
		delay   = 2     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	// this is also a composite literal - in this case a struct
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Ignoring encoding errors
}
