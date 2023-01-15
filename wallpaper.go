package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

type displayT struct {
	Name    string
	Rect    image.Rectangle
	PPI     int
	enabled bool
}

type renderFuncT func(rect image.Rectangle, display displayT) *image.NRGBA

type renderT struct {
	Name       string
	RenderFunc renderFuncT
	enabled    bool
}

const BASELINE_PPI = 109

func main() {
	displays := []displayT{
		{
			Name:    "m2_Air",
			Rect:    image.Rect(0, 0, 2560, 1664),
			PPI:     224,
			enabled: true,
		},
		{
			Name:    "MacBookPro_16",
			Rect:    image.Rect(0, 0, 3072, 1920),
			PPI:     226,
			enabled: true,
		},
		{
			Name:    "Dell_U4919DW",
			Rect:    image.Rect(0, 0, 5120, 1440),
			PPI:     109,
			enabled: true,
		},
	}

	renders := []renderT{
		{
			Name:       "hatch-saw-wave-blue",
			RenderFunc: createHatchSawWaveBlue,
			enabled:    true,
		},
		{
			Name:       "hatch-saw-wave-purple",
			RenderFunc: createHatchSawWavePurple,
			enabled:    true,
		},
		{
			Name:       "hatch-saw-wave-green",
			RenderFunc: createHatchSawWaveGreen,
			enabled:    true,
		},
		{
			Name:       "hatch-saw-wave-orange",
			RenderFunc: createHatchSawWaveOrange,
			enabled:    true,
		},
		{
			Name:       "hatch-saw-wave-teal",
			RenderFunc: createHatchSawWaveTeal,
			enabled:    true,
		},
		{
			Name:       "hatch-saw-blue",
			RenderFunc: createHatchSawBlue,
			enabled:    false,
		},
		{
			Name:       "hatch-saw-purple",
			RenderFunc: createHatchSawPurple,
			enabled:    false,
		},
		{
			Name:       "hatch-saw-green",
			RenderFunc: createHatchSawGreen,
			enabled:    false,
		},
		{
			Name:       "hatch-tri-blue",
			RenderFunc: createBlueHatchTri,
			enabled:    false,
		},
		{
			Name:       "hatch-sine-blue",
			RenderFunc: createBlueHatch,
			enabled:    false,
		},
		{
			Name:       "hatch-sine-purple",
			RenderFunc: createPurpleHatch,
			enabled:    false,
		},
		{
			Name:       "hatch-sine-green",
			RenderFunc: createGreenHatch,
			enabled:    false,
		},
		{
			Name:       "plumset",
			RenderFunc: createPlumset,
			enabled:    false,
		},
		{
			Name:       "melon",
			RenderFunc: createMelon,
			enabled:    false,
		},
		// {
		// 	Name:       "drops-blue",
		// 	RenderFunc: createBluedrops,
		// 	enabled:    false,
		// },
	}

	for _, display := range displays {
		display.showInfo()
		if !display.enabled {
			fmt.Println("   DISABLED")
			continue
		}
		var img *image.NRGBA
		rect := display.Rect
		var filename string

		for _, render := range renders {
			if !render.enabled {
				continue
			}
			filename = fmt.Sprintf(
				"img/wallpaper_%s_%dx%d__%s.png",
				display.Name,
				rect.Dx(),
				rect.Dy(),
				render.Name)
			fmt.Printf("   Creating %s...\n", filename)
			img = render.RenderFunc(rect, display)
			save(filename, img)
		}
	}
}

func distance(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(
		math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

func distanceF(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(
		math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func distanceFSq(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
}

func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}

func (d displayT) showInfo() {
	fmt.Printf(
		"%s: %d x %d @ %d PPI (%0.2fin x %0.2fin)\n",
		d.Name,
		d.Rect.Dx(),
		d.Rect.Dy(),
		d.PPI,
		float64(d.Rect.Dx())/float64(d.PPI),
		float64(d.Rect.Dy())/float64(d.PPI),
	)
}
