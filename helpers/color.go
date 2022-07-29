package helpers

import (
	"image/color"
	"math/rand"
	"time"
)

func RandomRGBAColor(maxR int, maxG int, maxB int, a int) color.RGBA {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return color.RGBA{R: uint8(r.Intn(maxR)), G: uint8(r.Intn(maxG)), B: uint8(r.Intn(maxB)), A: uint8(a)}
}
