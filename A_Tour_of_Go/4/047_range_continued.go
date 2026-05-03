package atourofgo

import "fmt"

func range_continued_test() {
	// Создаём слайс из 10 элементов, все изначально = 0
	pow := make([]int, 10)

	// Первый цикл: заполняем слайс степенями двойки
	for i := range pow { // здесь нам нужен только индекс
		pow[i] = 1 << uint(i) // 1 << i = 2^i
	}

	// Второй цикл: выводим значения
	for _, value := range pow { // здесь нам не нужен индекс, поэтому _
		fmt.Printf("%d\n", value)
	}
}
