package atourofgo

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // отправляем следующее число Фибоначчи
			x, y = y, x+y
		case <-quit: // получили сигнал остановки
			fmt.Println("quit")
			return
		}
	}
}

func select_qwe() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // читаем 10 чисел
		}
		quit <- 0 // отправляем сигнал завершения
	}()

	fibonacci(c, quit)
}
