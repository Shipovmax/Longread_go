package atourofgo

import "fmt"

func making_slices() {
	// 1. Создаём слайс длиной 5 и вместимостью 5
	a := make([]int, 5)
	printSlice_s("a", a)

	// 2. Создаём слайс длиной 0, но с вместимостью 5
	b := make([]int, 0, 5)
	printSlice_s("b", b)

	// 3. Берём срез из b (первые 2 элемента)
	c := b[:2]
	printSlice_s("c", c)

	// 4. Берём срез из c (с 3-го по 5-й элемент)
	d := c[2:5]
	printSlice_s("d", d)
}

func printSlice_s(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
