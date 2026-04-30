package atourofgo

import "fmt"

func deferrr() {
	defer fmt.Println("world")

	fmt.Println("Hello")
}
