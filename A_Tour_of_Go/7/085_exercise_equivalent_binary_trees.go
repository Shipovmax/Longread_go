package atourofgo

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk обходит дерево t в порядке inorder и отправляет все значения в канал ch.
// После обхода канал обязательно закрывается.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)  // левое поддерево
	ch <- t.Value     // текущее значение
	Walk(t.Right, ch) // правое поддерево
}

// Same проверяет, содержат ли два дерева одинаковые значения.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 && !ok2 { // оба канала закрыты
			return true
		}
	}
}

func exercise_equivalent_binary_trees() {
	// Тест
	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false
}
