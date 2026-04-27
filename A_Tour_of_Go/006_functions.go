package atourofgo

import "fmt"

// Инициализация каждой переменной
//func add(x int, y int) int {

// Инициализация группы переменных
func add(x, y int) int {
	return x + y
}

func test_add() {
	fmt.Println(add(42, 22))
}
