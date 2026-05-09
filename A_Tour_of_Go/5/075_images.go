package atourofgo

import (
	"fmt"
	"image"
)

func images() {
	// Создаём новое RGBA изображение размером 100x100
	// Новое изображение по умолчанию полностью чёрное и прозрачное
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Bounds возвращает прямоугольник изображения (границы)
	fmt.Println(m.Bounds())

	// At(x, y) возвращает цвет пикселя в точке (0, 0)
	// RGBA() возвращает 4 значения: R, G, B, A (каждое от 0 до 65535)
	fmt.Println(m.At(0, 0).RGBA())
}
