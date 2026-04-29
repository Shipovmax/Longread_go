package atourofgo

import "fmt"

func for_is_gos_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
