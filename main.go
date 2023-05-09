package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

// func rn(n int) int { return rand.Intn(n) }

func main() {
	f, err := os.OpenFile("demo.svg", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// data := []int{10, 33, 73, 64}
	// _ = data

	canvas := svg.New(f)
	data := []struct {
		Month string
		Usage int
	}{
		{"Jan", 171},
		{"Feb", 180},
		{"Mar", 100},
		{"Apr", 87},
		{"May", 66},
		{"Jun", 40},
		{"Jul", 32},
		{"Aug", 55},
		{"Sep", 63},
		{"Oct", 73},
		{"Nov", 83},
		{"Dec", 93},
	}
	width := len(data)*60 + 10
	height := 300
	threshold := 120
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	canvas.Start(width, height)
	for i, val := range data {
		percent := val.Usage * (height - 50) / max
		canvas.Text(i*60+35, height-20, val.Month, "font-family:Tahoma;font-size:14pt;fill:rgb(150,150,150);text-anchor:middle")
		canvas.Rect(i*60+10, (height-50)-percent, 50, percent, "fill:rgb(77,200,232)")
	}
	threshPercent := threshold * (height - 50) / max
	canvas.Line(0, height-threshPercent, width, height-threshPercent, "stroke: red;opacity:0.8;stroke-width:2px")
	canvas.Rect(0, 0, width, height-threshPercent, "fill:rgb(255,100,100);opacity:0.3")
	canvas.Line(0, height-50, width, height-50, "stroke: rgb(150, 150, 150); stroke-width:2")
	// nstars := 250
	// style := "font-size:48pt;fill:white;text-anchor:middle"
	// rand.Seed(time.Now().Unix())
	// canvas.Rect(0,0,width,height)
	// for i := 0; i < nstars; i++ {
	// 	canvas.Circle(rn(width), rn(height), rn(3), "fill:white")
	// }
	// canvas.Circle(width/2, height, width/2, "fill:rgb(77, 117, 232)")
	// canvas.Text(width/2, height*4/5, "hello, world", style)
	canvas.End()
}

// func main() {
// 	f, err := os.OpenFile("demo.svg", os.O_RDWR|os.O_CREATE, 0755)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	data := []int{10, 33, 73, 64}
// 	_ = data

// 	w, h := len(data)*60+10, 100
// 	canvas := svg.New(f)
// 	canvas.Start(w, h)
// 	canvas.Circle(w/2, h/2, 100)
// 	canvas.Text(w/2, h/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
// 	canvas.End()

// 	// r := image.Rect(0, 0, w, h)
// 	// img := image.NewRGBA(r)
// 	// grey := image.NewUniform(color.RGBA{240, 240, 240, 255})
// 	// blue := image.NewUniform(color.RGBA{180, 180, 250, 255})
// 	// green := image.NewUniform(color.RGBA{0, 255, 85, 255})

// 	// draw.Draw(img, r, grey, image.Point{0, 0}, draw.Src)

// 	// mask := image.NewNRGBA(image.Rect(0, 0, w, h))

// 	// for y := 0; y < h; y++ {
// 	// 	for x := 0; x < w; x++ {
// 	// 		mask.Set(x, y, color.NRGBA{
// 	// 			R: uint8((x + y) & 255),
// 	// 			G: uint8((x + y) << 1 & 255),
// 	// 			B: uint8((x + y) << 2 & 255),
// 	// 			A: 255,
// 	// 		})
// 	// 	}
// 	// }
// 	// for y := 0; y < h; y++ {
// 	// 	for x := 0; x < w; x++ {
// 	// 		img.Set(x, y, color.RGBA{255, 255, 255, 255})
// 	// 	}
// 	// }

// 	// for i, dp := range data {
// 	// 	x0, y0 := (i*60 + 10), 100-dp
// 	// 	x1, y1 := (i+1)*60-1, 100
// 	// 	bar := image.Rect(x0, y0, x1, y1)
// 	// 	// draw.Draw(img, bar, green, image.Point{0, 0}, draw.Src)
// 	// 	draw.Draw(img, bar, mask, image.Point{x0, y0}, draw.Src)

// 	// 	// for x := i*60 + 10; x < (i+1)*60; x++ {
// 	// 	// 	for y := 100; y >= (100 - dp); y-- {
// 	// 	// 		img.Set(x, y, color.RGBA{180, 180, 250, 255})
// 	// 	// 	}
// 	// 	// }
// 	// }

// 	// f, err := os.Create("image.png")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// defer f.Close()

// 	// err = png.Encode(f, img)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// }
