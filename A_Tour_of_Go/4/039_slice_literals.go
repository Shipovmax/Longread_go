package atourofgo

import "fmt"

func slice_literals() {
	// 1. Слайс целых чисел
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// 2. Слайс булевых значений
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// 3. Слайс анонимных структур
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
