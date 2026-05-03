package atourofgo

import "fmt"

func slides_pointers() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("Оригинальный массив:", names)

	// Создаём два слайса, которые смотрят на один и тот же массив
	a := names[0:2] // слайс от 0 до 2 (не включая 2) → [John, Paul]
	b := names[1:3] // слайс от 1 до 3 (не включая 3) → [Paul, George]

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Изменяем элемент через слайс b
	b[0] = "XXX"

	fmt.Println("\nПосле изменения b[0] = \"XXX\":")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("Оригинальный массив names =", names)
}
