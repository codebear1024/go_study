package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/echoRequest", echoRequest)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/showlissajous", showlissajous)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoRequest(w http.ResponseWriter, r *http.Request) {
	countInc()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
}

func countInc() {
	mu.Lock()
	count++
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	//countInc()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintln(w, "这是一个简单的web服务。")
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Fprintf(w, "Count = %d\n", count)
}

func showlissajous(w http.ResponseWriter, r *http.Request) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 400   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	// color.RGBA可得到不同的色值
	var colorRad = color.RGBA{R: 255, A: 0xff}
	var colorOrange = color.RGBA{R: 255, G: 97, A: 0xff}
	var colorYellow = color.RGBA{R: 255, G: 255, A: 0xff}
	var colorGreen = color.RGBA{G: 255, A: 0xff}
	var colorCyan = color.RGBA{G: 255, B: 255, A: 0xff}
	var colorBlue = color.RGBA{B: 255, A: 0xff}
	var colorPurple = color.RGBA{R: 160, G: 32, B: 240, A: 0xff}

	var palette = []color.Color{color.White, color.Black, colorRad, colorOrange, colorYellow, colorGreen, colorCyan, colorBlue, colorPurple}

	countInc()
	rand.Seed(time.Now().UTC().UnixNano())
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
				uint8(rand.Intn(8)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
