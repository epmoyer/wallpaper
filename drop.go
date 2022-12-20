package main

import "math"

type dropT struct {
	x         float64
	y         float64
	amplitude float64
	size      float64
}

type dropFieldT struct {
	Drops []dropT
}

func (drop dropT) render(x float64, y float64) (r, g, b float64) {
	d := distanceF(x, y, drop.x, drop.y)
	if d > drop.size {
		return 0, 0, 0
	}
	angle := d / 5.0
	depth := math.Sin(angle)
	depth *= (drop.size - d) / drop.size
	depth *= drop.amplitude
	return depth, depth, depth
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
