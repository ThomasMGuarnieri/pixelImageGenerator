package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"pixelImageGenerator/helpers"
	"strconv"
	"time"
)

var c color.RGBA

func main() {
	genStripedImage()
}

// genStripedImage generates a new striped image and save the file on root
func genStripedImage() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Set a random first color
	c = helpers.RandomRGBAColor(255, 255, 255, 255)

	m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	b := m.Bounds()

	// Loop through the image array, setting variants of the first color for every pixel
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			setRandomRGBAVariantColor(m, x, y, r)
		}
	}

	helpers.SavePNGImage(m, fmt.Sprintf("./stripedImages/%s.png", strconv.Itoa(r.Intn(255))))
}

func setRandomRGBAColor(m *image.RGBA, x int, y int) {
	m.SetRGBA(x, y, helpers.RandomRGBAColor(255, 255, 255, 255))
}

func setRandomRGBAVariantColor(m *image.RGBA, x int, y int, r *rand.Rand) {
	n := r.Intn(3)
	v := 2

	if n == 0 {
		c.R = getVariantValue(c.R, v, r)
	} else if n == 1 {
		c.G = getVariantValue(c.G, v, r)
	} else if n == 2 {
		c.B = getVariantValue(c.B, v, r)
	}

	m.SetRGBA(x, y, c)
}

// getVariantValue returns a new variant value from the given color, between 0 and 255
func getVariantValue(v uint8, sd int, r *rand.Rand) uint8 {
	b := r.Intn(2)
	usd := uint8(sd)

	if v > (255 - usd) {
		return v - uint8(r.Intn(sd))
	} else if v < usd {
		return v + uint8(r.Intn(sd))
	}

	if b == 1 {
		return v + uint8(r.Intn(sd))
	}

	return v - uint8(r.Intn(sd))
}
