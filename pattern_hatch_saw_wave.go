package main

import (
	"image"
	"math"

	"github.com/disintegration/imaging"
)

func createHatchSawWaveGreen(rect image.Rectangle, display displayT) *image.NRGBA {
	img := createHatchSawWaveBlue(rect, display)
	hue(img, -25)
	img = imaging.AdjustSaturation(img, -60)
	img = imaging.AdjustBrightness(img, -3)
	return img
}

func createHatchSawWavePurple(rect image.Rectangle, display displayT) *image.NRGBA {
	img := createHatchSawWaveBlue(rect, display)
	hue(img, +30)
	img = imaging.AdjustSaturation(img, -60)
	img = imaging.AdjustBrightness(img, 5)
	return img
}

func createHatchSawWaveOrange(rect image.Rectangle, display displayT) *image.NRGBA {
	img := createHatchSawWaveBlue(rect, display)
	hue(img, +180)
	img = imaging.AdjustSaturation(img, -10)
	// img = imaging.AdjustBrightness(img, -3)
	return img
}

func createHatchSawWaveTeal(rect image.Rectangle, display displayT) *image.NRGBA {
	img := createHatchSawWaveBlue(rect, display)
	hue(img, -35)
	img = imaging.AdjustSaturation(img, -30)
	img = imaging.AdjustBrightness(img, -6)
	return img
}

func createHatchSawWaveBlue(rect image.Rectangle, display displayT) *image.NRGBA {
	var r, g, b float64
	widthPix := rect.Dx()
	heightPix := rect.Dy()
	widthIn := float64(widthPix) * float64(display.PPI)
	pix := make([]uint8, widthPix*heightPix*4)
	stride := widthPix * 4

	for x := 0; x < widthPix; x++ {
		for y := 0; y < heightPix; y++ {

			yNormalized := float64(y) / float64(heightPix)
			xIn := float64(x) / float64(display.PPI)
			yIn := float64(y) / float64(display.PPI)

			// Wave
			angleWave := (float64)(xIn+yIn) * 6
			// angleWave *= float64(BASELINE_PPI) / float64(display.PPI)
			amplitudeWave := 0.03 * float64(display.PPI) * math.Sin(angleWave)
			// amplitudeWave *= float64(BASELINE_PPI) / float64(display.PPI)

			// Sawtooth
			angleSaw := (float64)(xIn+widthIn-yIn) * math.Pi * 2 * 6
			// angleSaw *= float64(BASELINE_PPI) / float64(display.PPI)
			angleSaw += amplitudeWave
			amplitude := -(math.Mod(angleSaw, 2*math.Pi) - math.Pi) / math.Pi

			base := x*4 + y*stride
			shade := (amplitude + 1) * 6
			fade := 80 * (1 - yNormalized)

			r = 50 + 10 + shade + fade
			g = 50 + 40 + shade*1.2 + fade
			b = 50 + 70 + shade + fade

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

	saturate(img, 0.1)
	img = imaging.Blur(img, 0.5)

	return img
}
