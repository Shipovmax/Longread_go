package atourofgo

import "fmt"

func arrays() {
	// Способ 1: объявление + присваивание
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Способ 2: короткое объявление (рекомендуется)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
