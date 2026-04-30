package atourofgo

import (
	"fmt"
	"time"
)

func switch_evaluation_order() {
	fmt.Println("When is Saturday?")

	// today — сегодняшний день недели (например, time.Monday, time.Wednesday и т.д.)
	today := time.Now().Weekday()

	// Мы переключаемся по константе time.Saturday (это 6)
	switch time.Saturday {

	// case'ы вычисляются по порядку сверху вниз
	case today + 0: // если сегодня суббота
		fmt.Println("Today")

	case today + 1: // если завтра суббота
		fmt.Println("Tomorrow.")

	case today + 2: // если послезавтра суббота
		fmt.Println("In two days.")

	default: // если до субботы больше 2 дней
		fmt.Println("Too far away.")
	}
}
