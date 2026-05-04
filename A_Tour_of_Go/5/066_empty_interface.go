package atourofgo

import "fmt"

// EmptyI — пустой интерфейс
type EmptyI interface {
	M()
}

// EmptyT — структура
type EmptyT struct {
	S string
}

// Метод для EmptyT
func (t *EmptyT) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// Функция вывода состояния
func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func empty_interface() {
	var i interface{} // пустой интерфейс

	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "hello"
	describeEmpty(i)

	i = &EmptyT{"World"}
	describeEmpty(i)
	i.(EmptyI).M() // вызов метода через type assertion
}
