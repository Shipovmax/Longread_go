package atourofgo

import "fmt"

// mx — map (словарь), где ключ — string, значение — Vertex
var mx = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967}, // можно опускать Vertex{}
	"Google":    {37.42202, -122.08408},
}

func map_literals() {
	fmt.Println(mx)
}
