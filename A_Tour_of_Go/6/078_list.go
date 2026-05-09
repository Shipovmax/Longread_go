package atourofgo

import "fmt"

// List — односвязный список с generics
// T any — может хранить значения любого типа
type List[T any] struct {
	next *List[T]
	val  T
}

// Push добавляет новый элемент в конец списка
func (l *List[T]) Push(v T) {
	if l.next == nil {
		l.next = &List[T]{val: v}
		return
	}
	l.next.Push(v)
}

// String для красивого вывода списка
func (l *List[T]) String() string {
	if l == nil {
		return "[]"
	}
	s := fmt.Sprintf("%v", l.val)
	if l.next != nil {
		s += " -> " + l.next.String()
	}
	return s
}

func list() {
	var lst List[int]

	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	fmt.Println(lst) // 10 -> 20 -> 30
}
