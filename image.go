package main

import (
	"crypto/rand"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 2560, 1664)
	var img *image.NRGBA
	// img = createTestImage1(rect)
	// save("image1.png", img)
	img = createTestImage2(rect)
	save("image2.png", img)
}

func createTestImage2(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	stride := rect.Dx() * 4
	m1X := rect.Dx()
	// m1X := 0
	m1Y := 0
	diagonal := math.Sqrt(float64(rect.Dx()*rect.Dx() + rect.Dy()*rect.Dy()))
	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			m2Distance := distance(x, y, m1X, m1Y)
			m2DistanceNorm := m2Distance / diagonal
			yNorm := float64(y) / float64(rect.Dy())
			angle := (float64)(x+y) / 5.4567
			// angle2 := (float64)(x+y) / 20.199782
			amplitude := math.Sin(angle)
			// amplitude2 := math.Sin(angle2)
			base := x*4 + y*stride
			shade := uint8((amplitude+1)*5) + 200
			pix[base] = shade - 100 + uint8(120*yNorm)
			// pix[base+1] = shade - 50 + uint8((amplitude2+1)*10)
			var bump uint8
			// bump = 0
			// if m2DistanceNorm < 0.5 {
			// 	// if m2Distance == -100 {
			// 	bump = uint8(30 * (1 - m2DistanceNorm))
			// }
			bump = uint8(20 * (1 - m2DistanceNorm))
			pix[base+1] = shade - 50
			// pix[base+2] = shade + 30 + uint8(20*m2DistanceNorm)
			pix[base+2] = shade + 30 + bump
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
