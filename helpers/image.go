package helpers

import (
	"image"
	"image/png"
	"os"
)

func SavePNGImage(m *image.RGBA, fileName string) {
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, m.SubImage(image.Rect(0, 0, 640, 480)))
	if err != nil {
		panic(err)
	}
}
