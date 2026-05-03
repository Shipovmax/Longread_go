package atourofgo

import "fmt"

func mutating_maps() {
	// Создаём пустой map: ключ - string, значение - int
	m := make(map[string]int)

	// Добавляем пару ключ-значение
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) // 42

	// Перезаписываем значение по тому же ключу
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) // 48

	// Удаляем ключ из map
	delete(m, "Answer")

	// Попытка прочитать несуществующий ключ
	fmt.Println("The value:", m["Answer"]) // 0 (нулевое значение для int)

	// Самый правильный способ проверки наличия ключа
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // 0 false
}
