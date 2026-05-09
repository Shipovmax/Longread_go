package atourofgo

import "fmt"

// Index — generic функция
// T comparable — значит T можно сравнивать через ==
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func index() {
	// Работа с int
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Работа с string
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
