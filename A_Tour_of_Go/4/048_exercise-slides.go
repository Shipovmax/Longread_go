package atourofgo

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Создаём двумерный слайс (dy строк, каждая длиной dx)
	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			// Здесь можно экспериментировать с разными формулами
			// Самый красивый и популярный вариант:
			pic[y][x] = uint8(x * y) // вариант 1
			// pic[y][x] = uint8(x ^ y)        // вариант 2 (XOR)
			// pic[y][x] = uint8(x + y)        // вариант 3
			// pic[y][x] = uint8(x * x + y * y) // вариант 4
		}
	}

	return pic
}

func exercise_slides() {
	pic.Show(Pic)
}
