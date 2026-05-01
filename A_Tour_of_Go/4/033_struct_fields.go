package atourofgo

import "fmt"

type Vertexs struct {
	X int
	Y int
}

func struct_fields() {
	v := Vertexs{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
