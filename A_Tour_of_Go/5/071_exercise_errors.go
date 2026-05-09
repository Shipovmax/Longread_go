package atourofgo

import "fmt"

// Sqrt — функция, которая вычисляет квадратный корень.
func Sqrt(x float64) (float64, error) {
	return 0, nil // никогда не возвращаем только результат,если может возникнуть ошибка.
}

// Вывод \ Тест
func exercise_errors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
