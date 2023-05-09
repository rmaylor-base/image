package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	data := []int{10, 33, 73, 64}
	w, h := len(data)*60+10, 100
	r := image.Rect(0, 0, w, h)
	img := image.NewRGBA(r)
	grey := image.NewUniform(color.RGBA{240, 240, 240, 255})
	// blue := image.NewUniform(color.RGBA{180, 180, 250, 255})
	// green := image.NewUniform(color.RGBA{0, 255, 85, 255})

	draw.Draw(img, r, grey, image.Point{0, 0}, draw.Src)

	mask := image.NewNRGBA(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			mask.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}
	// test
	// for y := 0; y < h; y++ {
	// 	for x := 0; x < w; x++ {
	// 		img.Set(x, y, color.RGBA{255, 255, 255, 255})
	// 	}
	// }

	for i, dp := range data {
		x0, y0 := (i*60+10), 100-dp
		x1, y1 := (i+1)*60-1, 100
		bar := image.Rect(x0, y0, x1, y1)
		// draw.Draw(img, bar, green, image.Point{0, 0}, draw.Src)
		draw.Draw(img, bar, mask, image.Point{x0, y0}, draw.Src)

		// for x := i*60 + 10; x < (i+1)*60; x++ {
		// 	for y := 100; y >= (100 - dp); y-- {
		// 		img.Set(x, y, color.RGBA{180, 180, 250, 255})
		// 	}
		// }
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
