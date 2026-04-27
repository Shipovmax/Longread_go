package atourofgo

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func return_swap() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
