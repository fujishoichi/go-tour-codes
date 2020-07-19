package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	pixels [][]uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	w := len(i.pixels)
	h := len(i.pixels[0])
	return image.Rect(0, 0, w, h)
}

func (i Image) At(x, y int) color.Color {
	pixel := i.pixels[x][y]
	return color.RGBA{pixel, pixel, 255, 255}
}

func Pixels(dx, dy int )[][]uint8 {
	pixels := make([][]uint8, dx)
	for x := range pixels {
		pixels[x] = make([]uint8, dy)
		for y := range pixels[x] {
			pixels[x][y] = uint8(x * y)
			//uint8(x^y)
			//uint8(x%(y+1))
		}
	}
	return pixels
}

func main() {
	m := Image{Pixels(200, 100)}
	pic.ShowImage(m)
}
