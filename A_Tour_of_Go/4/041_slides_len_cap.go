package atourofgo

import "fmt"

func slides_len_cap() {
	// Создаём слайс
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 1. Обрезаем слайс до нулевой длины (но вместимость остаётся!)
	s = s[:0]
	printSlice(s)

	// 2. Расширяем длину слайса (используем уже зарезервированную память)
	s = s[:4]
	printSlice(s)

	// 3. Отбрасываем первые два элемента
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
