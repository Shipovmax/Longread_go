package atourofgo

import "fmt"

func append_test() {
	// Создаём nil-слайс (самый обычный способ)
	var s []int
	printSlice_1(s)

	// append работает даже с nil-слайсом!
	s = append(s, 0)
	printSlice_1(s)

	// Добавляем ещё один элемент
	s = append(s, 1)
	printSlice_1(s)

	// Можно добавлять сразу несколько элементов
	s = append(s, 2, 3, 4)
	printSlice_1(s)
}

func printSlice_1(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
