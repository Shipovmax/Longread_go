package atourofgo

import "fmt"

// m — map (словарь), где ключ — название, значение — координаты
var mq = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func map_literals_continued() {
	fmt.Println(mq)
}
