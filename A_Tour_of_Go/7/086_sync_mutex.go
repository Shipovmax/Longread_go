package atourofgo

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter — безопасный счётчик для использования из нескольких горутин
type SafeCounter struct {
	mu sync.Mutex // мьютекс защищает доступ к map
	v  map[string]int
}

// Inc увеличивает значение по ключу
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // блокируем — только одна горутина может зайти
	c.v[key]++
	c.mu.Unlock() // разблокируем
}

// Value возвращает текущее значение по ключу
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // автоматически разблокируем при выходе
	return c.v[key]
}

func sync_mutex() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey")) // должно быть 1000
}
