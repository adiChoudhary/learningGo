package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (p Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 50, 50)
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 255, 255}
}

func main() {
	m := Image{}
	// Works only in Go Playground
	pic.ShowImage(m)
}
