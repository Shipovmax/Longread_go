package atourofgo

import "fmt"

// I — интерфейс с одним методом
type I interface {
	M()
}

// T — обычная структура
type T struct {
	S string
}

// Метод с pointer receiver
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func indirection_values_with_nil() {
	var i I  // интерфейс
	var t *T // указатель на T (сейчас nil)

	// 1. Присваиваем nil-указатель в интерфейс
	i = t
	describe(i) // ( <nil> , *main.T )
	i.M()       // вызов метода на nil-указателе

	// 2. Присваиваем нормальное значение
	i = &T{"hello"}
	describe(i) // ( &{hello} , *main.T )
	i.M()       // hello
}

// Вспомогательная функция — показывает тип и значение интерфейса
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
