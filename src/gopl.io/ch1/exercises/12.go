package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var pallete = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const res = 0.001
	var cycles float64
	var size, nframes, delay int

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	rCycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		cycles = 5
	} else {
		cycles = float64(rCycles)
	}

	rFrames, err := strconv.Atoi(r.Form.Get("frames"))
	if err != nil {
		nframes = 64
	} else {
		nframes = rFrames
	}

	rsize, err := strconv.Atoi(r.Form.Get("size"))
	if err != nil {
		size = 100
	} else {
		size = rsize
	}

	rdelay, err := strconv.Atoi(r.Form.Get("delay"))
	if err != nil {
		delay = 8
	} else {
		delay = rdelay
	}

	lissajous(w, res, cycles, size, nframes, delay)
}

func lissajous(out io.Writer, res, cycles float64, size, nframes, delay int) {

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size-1, 2*size-1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
