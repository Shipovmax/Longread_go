package atourofgo

import "fmt"

// Person — обычная структура
type Person struct {
	Name string
	Age  int
}

// String() — метод, который реализует интерфейс fmt.Stringer
// Теперь fmt.Println будет автоматически использовать этот метод
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println(a, z)
}
