package atourofgo

import (
	"fmt"
	"time"
)

func switch_with_no_condition() {
	t := time.Now()

	// switch без выражения — каждый case проверяет своё условие
	switch {
	case t.Hour() < 12: // если сейчас до 12:00
		fmt.Println("Good morning")

	case t.Hour() < 17: // если сейчас от 12:00 до 17:00
		fmt.Println("Good afternoon")

	default: // после 17:00
		fmt.Println("Good evening")
	}
}
