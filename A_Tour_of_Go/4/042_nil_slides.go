package atourofgo

import "fmt"

func nil_slides() {
	var s []int // ← вот так объявляется nil-слайс
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}
