package atourofgo

import "fmt"

func short_variable_declarations() {

	// Явное обозначение
	var i, j int = 1, 2

	// Самое короткое только внутри функции
	k := 3

	// Обьявлять только внутри функции и новая переменная слева обязательная
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
