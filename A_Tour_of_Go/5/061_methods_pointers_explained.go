package atourofgo

import (
	"fmt"
)

// Scale — обычная функция, которая принимает указатель
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methods_pointers_explained() {
	v := Vertex{3, 4}

	// Передаём указатель на v
	Scale(&v, 10)

	// Передаём копию v
	fmt.Println(Abs(v)) // → 50
}
