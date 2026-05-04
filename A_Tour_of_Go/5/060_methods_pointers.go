package atourofgo

import (
	"fmt"
)

// Scale — метод с pointer receiver (изменяет оригинальную структуру)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methods_pointers() {
	v := Vertex{3, 4}

	v.Scale(10)          // изменяем оригинал
	fmt.Println(v.Abs()) // должно вывести 50
}
