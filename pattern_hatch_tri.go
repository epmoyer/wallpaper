package main

import (
	"image"
	"math"

	"github.com/disintegration/imaging"
)

func createBlueHatchTri(rect image.Rectangle, display displayT) *image.NRGBA {
	var r, g, b float64
	width := rect.Dx()
	height := rect.Dy()
	pix := make([]uint8, width*height*4)
	stride := width * 4
	// m1X := width / 2
	m1X := 0
	m1Y := 0
	diagonal := math.Sqrt(float64(width*width + height*height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			m2Distance := distance(x, y, m1X, m1Y)
			m2DistanceNorm := m2Distance / diagonal
			bump := 72 * (1 - m2DistanceNorm)
			yNorm := float64(y) / float64(height)
			angle := (float64)(x+width-y) / 2.12345
			angle *= float64(BASELINE_PPI) / float64(display.PPI)
			// angle /= 10
			// amplitude := math.Sin(angle)

			// Triangle
			amplitude := (math.Abs(math.Mod(angle, 2*math.Pi)-math.Pi) - (math.Pi / 2)) / (math.Pi / 2)

			// Sawtooth
			// amplitude := (math.Mod(angle, 2*math.Pi) - math.Pi) / math.Pi

			base := x*4 + y*stride
			shade := (amplitude + 1) * 10
			fade := 80 * (1 - yNorm)

			r = 50 + 10 + shade + fade
			g = 50 + 40 + shade*1.2 + fade
			b = 50 + 70 + shade + fade + 0.00001*bump

			pix[base+0] = uint8(r)
			pix[base+1] = uint8(g)
			pix[base+2] = uint8(b)
			pix[base+3] = 255
		}
	}
	img := &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	// brighten(img, 0.1)
	// lighten(created, 0.03)
	saturate(img, 0.1)
	// hue(img, -55)
	// created =
	// img = imaging.Blur(img, 3.2)
	img = imaging.AdjustBrightness(img, -20)
	return img
}
