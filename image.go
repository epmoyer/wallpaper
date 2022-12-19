package main

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/disintegration/imaging"
	"github.com/teacat/noire"
)

type displayT struct {
	Name string
	Rect image.Rectangle
}

func main() {
	displays := []displayT{
		{
			Name: "m2_Air",
			Rect: image.Rect(0, 0, 2560, 1664),
		},
		// {
		// 	Name: "MacBookPro_16",
		// 	Rect: image.Rect(0, 0, 3072, 1920),
		// },
		// {
		// 	Name: "Dell_U4919DW",
		// 	Rect: image.Rect(0, 0, 5120, 1440),
		// },
	}
	for _, display := range displays {
		var img *image.NRGBA
		rect := display.Rect

		// filename := fmt.Sprintf("wallpaper_plumset_%s_%dx%d.png", display.Name, rect.Dx(), rect.Dy())
		// fmt.Printf("Creating %s...\n", filename)
		// img = createPlumset(rect)
		// save(filename, img)

		filename := fmt.Sprintf("wallpaper_WIP_%s_%dx%d.png", display.Name, rect.Dx(), rect.Dy())
		fmt.Printf("Creating %s...\n", filename)
		img = createWIP(rect)
		save(filename, img)
	}
}

func createWIP(rect image.Rectangle) *image.NRGBA {
	var r, g, b float64
	width := rect.Dx()
	height := rect.Dy()
	pix := make([]uint8, width*height*4)
	stride := width * 4
	m1X := width / 2
	m1Y := 0
	diagonal := math.Sqrt(float64(width*width + height*height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			m2Distance := distance(x, y, m1X, m1Y)
			m2DistanceNorm := m2Distance / diagonal
			bump := 72 * (1 - m2DistanceNorm)
			yNorm := float64(y) / float64(height)
			angle := (float64)(x+width-y) / 5.4567
			amplitude := math.Sin(angle)
			base := x*4 + y*stride
			shade := (amplitude + 1) * 5
			fade := 70 * yNorm

			r = 128 + 10 + shade
			g = 128 + shade*1.2 + fade
			b = 128 + 30 + shade + bump

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
	// brighten(img, 0.23)
	// lighten(created, 0.03)
	// saturate(img, 0.25)
	// hue(img, -55)
	// created =
	// img = imaging.Blur(img, 3.2)
	// img = imaging.(img, 3.2)
	return img
}

func createPlumset(rect image.Rectangle) *image.NRGBA {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	stride := rect.Dx() * 4
	m1X := rect.Dx()
	m1Y := 0
	diagonal := math.Sqrt(float64(rect.Dx()*rect.Dx() + rect.Dy()*rect.Dy()))
	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			m2Distance := distance(x, y, m1X, m1Y)
			m2DistanceNorm := m2Distance / diagonal
			yNorm := float64(y) / float64(rect.Dy())
			angle := (float64)(x+y) / 5.4567
			amplitude := math.Sin(angle)
			base := x*4 + y*stride
			shade := uint8((amplitude+1)*5) + 150
			pix[base] = shade - 100 + uint8(120*yNorm)
			bump := uint8(72 * (1 - m2DistanceNorm))
			pix[base+1] = shade - 40 + bump
			pix[base+2] = shade + 20 + bump
			pix[base+3] = 255
		}
	}
	img := &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	brighten(img, 0.23)
	// lighten(created, 0.03)
	saturate(img, 0.25)
	// created =
	img = imaging.Blur(img, 3.2)
	return img
}

func createTestImage4(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	stride := rect.Dx() * 4
	m1X := rect.Dx()
	m1Y := 0
	diagonal := math.Sqrt(float64(rect.Dx()*rect.Dx() + rect.Dy()*rect.Dy()))
	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			m2Distance := distance(x, y, m1X, m1Y)
			m2DistanceNorm := m2Distance / diagonal
			yNorm := float64(y) / float64(rect.Dy())
			angle := (float64)(x+y) / 5.4567
			// angle := (float64)(x+y) / 10.4567
			// angle2 := (float64)(x+y) / 20.199782
			amplitude := math.Sin(angle)
			// amplitude2 := math.Sin(angle2)
			base := x*4 + y*stride
			shade := uint8((amplitude+1)*5) + 150
			// pix[base+1] = shade - 50 + uint8((amplitude2+1)*10)
			var bump uint8
			// bump = 0
			// if m2DistanceNorm < 0.5 {
			// 	// if m2Distance == -100 {
			// 	bump = uint8(30 * (1 - m2DistanceNorm))
			// }
			bump = uint8(75 * (1 - m2DistanceNorm))
			pix[base+0] = shade + 0
			pix[base+1] = shade - 100 + uint8(120*yNorm)
			// pix[base+2] = shade - 100 + uint8(120*yNorm) + bump
			pix[base+2] = shade + 15 + bump
			pix[base+3] = 255
			// fmt.Println(x)
		}
	}
	created = &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}

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

func createTestImage3(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	stride := rect.Dx() * 4
	// m1X := rect.Dx()
	m1X := 0
	m1Y := rect.Dy()
	diagonal := math.Sqrt(float64(rect.Dx()*rect.Dx() + rect.Dy()*rect.Dy()))
	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			m2Distance := distance(x, y, m1X, m1Y)
			m2DistanceNorm := m2Distance / diagonal
			yNorm := float64(y) / float64(rect.Dy())
			angle := (float64)(x+y) / 5.4567
			// angle := (float64)(x+y) / 10.4567
			// angle2 := (float64)(x+y) / 20.199782
			amplitude := math.Sin(angle)
			// amplitude2 := math.Sin(angle2)
			base := x*4 + y*stride
			shade := uint8((amplitude+1)*5) + 150
			pix[base] = shade - 100 + uint8(120*yNorm)
			// pix[base+1] = shade - 50 + uint8((amplitude2+1)*10)
			var bump uint8
			// bump = 0
			// if m2DistanceNorm < 0.5 {
			// 	// if m2Distance == -100 {
			// 	bump = uint8(30 * (1 - m2DistanceNorm))
			// }
			bump = uint8(75 * (1 - m2DistanceNorm))
			pix[base+0] = shade - 40 + bump
			// pix[base+2] = shade + 30 + uint8(20*m2DistanceNorm)
			pix[base+1] = shade + 20 + bump
			pix[base+2] = 255
			pix[base+3] = 255
			// fmt.Println(x)
		}
	}
	created = &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}

func distance(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(
		math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

func createTestImage1(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	stride := rect.Dx() * 4
	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			yNorm := float64(y) / float64(rect.Dy())
			angle := (float64)(x+y) / 5.4567
			amplitude := math.Sin(angle)
			base := x*4 + y*stride
			shade := uint8((amplitude+1)*5) + 200
			pix[base] = shade - 100 + uint8(120*yNorm)
			pix[base+1] = shade - 50
			pix[base+2] = shade + 30
			pix[base+3] = 255
			// fmt.Println(x)
		}
	}
	// rand.Read(pix)
	created = &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}

func createRandomImage(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	rand.Read(pix)
	created = &image.NRGBA{
		Pix:    pix,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}

func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	png.Encode(imgFile, img.SubImage(img.Rect))
}
