package atourofgo

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

// Bounds возвращает размер изображения
func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 300, 300)
}

// ColorModel возвращает цветовую модель
func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At возвращает цвет пикселя в точке (x, y)
func (Image) At(x, y int) color.Color {
	// Красивый паттерн (XOR эффект)
	r := uint8(x ^ y)
	g := uint8(x + y)
	b := uint8(x * y)

	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func exercise_images() {
	m := Image{}
	pic.ShowImage(m)
}
