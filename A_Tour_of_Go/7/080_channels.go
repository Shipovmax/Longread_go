package atourofgo

import "fmt"

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total // отправляем результат в канал
}

func channels() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // создаём небуферизированный канал

	// запускаем две горутины, каждая считает свою половину
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// получаем результаты из канала (в любом порядке)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
