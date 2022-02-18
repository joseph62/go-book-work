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

var red = color.RGBA{R: 0xff, A: 0xff}
var orange = color.RGBA{R: 0xff, G: 0xa5, A: 0xff}
var yellow = color.RGBA{R: 0xff, G: 0xff, A: 0xff}
var green = color.RGBA{G: 0x80, A: 0xff}
var blue = color.RGBA{B: 0xff, A: 0xff}
var indigo = color.RGBA{R: 0x4b, B: 0x82, A: 0xff}
var violet = color.RGBA{R: 0xee, G: 0x82, B: 0xee, A: 0xff}

var pallete = []color.Color{color.White, red, orange, yellow, green, blue, indigo, violet}

func makeNumberIterator(numbers []uint8, repetitions uint) func() uint8 {
	var currentCount, currentNumber uint
	return func() uint8 {
		if currentCount > repetitions {
			currentCount = 0
			currentNumber = (currentNumber + 1) % uint(len(numbers))
		} else {
			currentCount++
		}
		return numbers[currentNumber]
	}
}

var rainbowIterator = makeNumberIterator([]uint8{1, 2, 3, 4, 5, 6, 7}, 1024)

const (
	backgroundIndex = 0
	foregroundIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 8
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size-1, 2*size-1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), rainbowIterator())
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
