package main

import "math"

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
