package atourofgo

import "fmt"

// Vertexqweasd — структура с координатами
type Vertex struct {
	Lat, Long float64
}

// m — объявляем переменную типа map
// Ключ: string, Значение: Vertex
var m map[string]Vertex

func maps() {
	// Важно! Map нужно инициализировать через make
	m = make(map[string]Vertex)

	// Добавляем элемент в map
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}
