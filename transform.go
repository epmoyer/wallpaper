package main

import (
	"image"

	"github.com/teacat/noire"
)

func lighten(image *image.NRGBA, lightenPercent float64) {
	var r, g, b float64
	for x := 0; x < image.Rect.Dx(); x++ {
		for y := 0; y < image.Rect.Dy(); y++ {
			base := x*4 + y*image.Stride
			c := noire.NewRGB(
				float64(image.Pix[base+0]),
				float64(image.Pix[base+1]),
				float64(image.Pix[base+2]))
			r, g, b = c.Lighten(lightenPercent).RGB()
			image.Pix[base+0] = uint8(r)
			image.Pix[base+1] = uint8(g)
			image.Pix[base+2] = uint8(b)
		}
	}
}

func brighten(image *image.NRGBA, percent float64) {
	var r, g, b float64
	for x := 0; x < image.Rect.Dx(); x++ {
		for y := 0; y < image.Rect.Dy(); y++ {
			base := x*4 + y*image.Stride
			c := noire.NewRGB(
				float64(image.Pix[base+0]),
				float64(image.Pix[base+1]),
				float64(image.Pix[base+2]))
			r, g, b = c.Brighten(percent).RGB()
			image.Pix[base+0] = uint8(r)
			image.Pix[base+1] = uint8(g)
			image.Pix[base+2] = uint8(b)
		}
	}
}

func saturate(image *image.NRGBA, saturatePercent float64) {
	var r, g, b float64
	for x := 0; x < image.Rect.Dx(); x++ {
		for y := 0; y < image.Rect.Dy(); y++ {
			base := x*4 + y*image.Stride
			c := noire.NewRGB(
				float64(image.Pix[base+0]),
				float64(image.Pix[base+1]),
				float64(image.Pix[base+2]))
			r, g, b = c.Saturate(saturatePercent).RGB()
			image.Pix[base+0] = uint8(r)
			image.Pix[base+1] = uint8(g)
			image.Pix[base+2] = uint8(b)
		}
	}
}

func hue(image *image.NRGBA, degrees float64) {
	var r, g, b float64
	for x := 0; x < image.Rect.Dx(); x++ {
		for y := 0; y < image.Rect.Dy(); y++ {
			base := x*4 + y*image.Stride
			c := noire.NewRGB(
				float64(image.Pix[base+0]),
				float64(image.Pix[base+1]),
				float64(image.Pix[base+2]))
			r, g, b = c.AdjustHue(degrees).RGB()
			image.Pix[base+0] = uint8(r)
			image.Pix[base+1] = uint8(g)
			image.Pix[base+2] = uint8(b)
		}
	}
}
