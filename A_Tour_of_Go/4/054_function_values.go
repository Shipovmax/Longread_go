package atourofgo

import (
	"fmt"
	"math"
)

// compute принимает функцию в качестве аргумента
// fn — это функция, которая принимает два float64 и возвращает float64
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4) // вызываем переданную функцию с аргументами 3 и 4
}

func function_values() {
	// Определяем анонимную функцию и сохраняем её в переменную hypot
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// 1. Вызываем функцию напрямую
	fmt.Println(hypot(5, 12)) // → 13  (гипотенуза треугольника 5-12-13)

	// 2. Передаём функцию hypot в compute
	fmt.Println(compute(hypot)) // → 5   (вычисляет hypot(3,4))

	// 3. Передаём встроенную функцию math.Pow
	fmt.Println(compute(math.Pow)) // → 81  (3^4 = 81)
}
