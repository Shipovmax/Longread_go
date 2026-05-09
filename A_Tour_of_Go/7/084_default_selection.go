package atourofgo

import (
	"fmt"
	"time"
)

func default_selection() {
	start := time.Now()

	tick := time.Tick(100 * time.Millisecond)  // тикер каждые 100мс
	boom := time.After(500 * time.Millisecond) // сработает через 500мс

	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	for {
		select {
		case <-tick: // сработал тикер
			fmt.Printf("[%6s] tick.\n", elapsed())

		case <-boom: // сработал таймер
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return

		default: // ни один канал не готов
			fmt.Printf("[%6s] .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
