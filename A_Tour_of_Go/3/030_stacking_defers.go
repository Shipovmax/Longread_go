package atourofgo

import "fmt"

func stacking_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // каждый defer откладывается
		// На 10-й итерации стек defer'ов выглядит так:
		// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]  ← 9 сверху
	}

	fmt.Println("done")

	// Здесь функция заканчивается → defers выполняются в обратном порядке
}
