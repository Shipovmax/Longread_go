package atourofgo

import "fmt"

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Создаём слайс из массива primes
	// primes[1:4] → элементы с индекса 1 до (но не включая) 4
	var s []int = primes[1:4]

	fmt.Println(s)      // [3 5 7]
	fmt.Println(len(s)) // 3  (длина)
	fmt.Println(cap(s)) // 5  (вместимость до конца массива)
}
