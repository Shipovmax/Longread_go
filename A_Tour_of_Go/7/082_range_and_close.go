package atourofgo

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x        // отправляем очередное число Фибоначчи
		x, y = y, x+y // вычисляем следующее
	}
	close(c) // обязательно закрываем канал после завершения
}

func range_and_close() {
	c := make(chan int, 10) // буферизированный канал на 10 элементов

	go fibonacci(cap(c), c) // запускаем генерацию в горутине

	// range по каналу автоматически читает до close(c)
	for i := range c {
		fmt.Println(i)
	}
}
