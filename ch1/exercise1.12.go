package main

import (
	"net/http"
	"log"
	"io"
	"image/gif"
	"image"
	"math/rand"
	"math"
	"image/color"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const blackIndex = 1

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, err := strconv.ParseFloat(r.FormValue("cycles"), 64)
	if err != nil {
		cycles = 5
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles float64) {
	if cycles <= 5 {
		cycles = 5
	}
	const (
		res 	= 0.001
		size 	= 100
		nframes = 64
		delay 	= 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
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
	gif.EncodeAll(out, &anim)
}
