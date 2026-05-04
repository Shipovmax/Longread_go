package atourofgo

import (
	"fmt"
	"math"
)

// Обычная функция (не метод)
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func indirection_values() {
	v := Vertex{3, 4}
	p := &Vertex{4, 3}

	fmt.Println("Метод на значении:     ", v.Abs())
	fmt.Println("Функция на значении:   ", AbsFunc(v))

	fmt.Println("Метод на указателе:    ", p.Abs())     // Go сам разыменовывает
	fmt.Println("Функция на указателе:  ", AbsFunc(*p)) // нужно явно разыменовать
}
