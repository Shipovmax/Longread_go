package atourofgo

import (
	"fmt"
	"math"
)

// Обычная функция (не метод)
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods_funcs() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // 5
}
