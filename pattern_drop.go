package main

import (
	"image"
	"math"
	"math/rand"
)

const lightAngle = 3 * math.Pi / 8

type dropT struct {
	x         float64
	y         float64
	amplitude float64
	size      float64
	sizeSq    float64
	waveSize  float64
}

type dropFieldT struct {
	Drops []dropT
}

func (drop *dropT) init() {
	drop.sizeSq = drop.size * drop.size
}

func (drop dropT) render(x float64, y float64) (r, g, b float64) {
	dsq := distanceFSq(x, y, drop.x, drop.y)
	if dsq > drop.sizeSq {
		// Out of range
		return 0, 0, 0
	}
	d := math.Sqrt(dsq)

	angleToDrop := math.Atan2(x-drop.x, y-drop.y) - lightAngle
	waveAngle := d / drop.waveSize
	depth := math.Sin(waveAngle)
	depth *= (drop.size - d) / drop.size
	depth *= drop.amplitude
	depth *= math.Sin(angleToDrop)
	return depth, depth / 4, depth / 4
}

func (dropField dropFieldT) init() {

	for i := 0; i < len(dropField.Drops); i++ {
		dropField.Drops[i].init()
	}
}

func (dropField dropFieldT) render(x float64, y float64) (r, g, b float64) {

	for i := 0; i < len(dropField.Drops); i++ {
		dr, dg, db := dropField.Drops[i].render(x, y)
		r += dr
		g += dg
		b += db
	}
	return r, g, b
}

func createBluedrops(rect image.Rectangle) *image.NRGBA {
	var r, g, b float64
	width := rect.Dx()
	height := rect.Dy()
	pix := make([]uint8, width*height*4)
	stride := width * 4
	// m1X := width / 2
	// m1X := 0
	// m1Y := 0
	// diagonal := math.Sqrt(float64(width*width + height*height))

	field := dropFieldT{}
	// field := dropFieldT{
	// 	Drops: []dropT{
	// 		{
	// 			x:         400.0,
	// 			y:         400.0,
	// 			amplitude: 30.0,
	// 			size:      200,
	// 			waveSize:  5,
	// 		},
	// 	},
	// }

	rand.Seed(1)
	for i := 0; i < 130; i++ {
		field.Drops = append(field.Drops, dropT{
			x:         rand.Float64() * float64(width),
			y:         rand.Float64() * float64(height),
			amplitude: 10 + rand.Float64()*30,
			size:      200 + 200*rand.Float64(),
			waveSize:  5 + rand.Float64()*5,
		})
	}
	field.init()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			// m2Distance := distance(x, y, m1X, m1Y)
			// m2DistanceNorm := m2Distance / diagonal
			// bump := 72 * (1 - m2DistanceNorm)
			yNorm := float64(y) / float64(height)
			// angle := (float64)(x+width-y) / 5.4567
			// amplitude := math.Sin(angle)
			// shade := (amplitude + 1) * 5
			fade := 64 * yNorm

			dr, dg, db := field.render(float64(x), float64(y))
			r = 70 + dr
			g = 70 + fade + dg
			b = 128 + 64 + db

			base := x*4 + y*stride
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
	// saturate(img, 0.1)
	// hue(img, -55)
	// created =
	// img = imaging.Blur(img, 3.2)
	// img = imaging.AdjustBrightness(img, -20)
	return img
}
