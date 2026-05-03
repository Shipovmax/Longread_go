package atourofgo

import "fmt"

// fibonacci возвращает функцию, которая при каждом вызове выдаёт следующее число Фибоначчи
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b // сдвигаем последовательность
		return a
	}
}

func exercise_fibonacci_closure() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
