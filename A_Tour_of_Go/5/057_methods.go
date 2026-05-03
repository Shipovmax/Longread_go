package atourofgo

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Метод Abs у структуры Vertex
// (v Vertex) — это "receiver" (получатель)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // → 5
}