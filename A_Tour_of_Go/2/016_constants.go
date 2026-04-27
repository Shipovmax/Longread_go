package atourofgo

import "fmt"

const Pi = 3.14

func constants() {
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
