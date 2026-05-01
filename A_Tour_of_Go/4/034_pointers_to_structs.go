package atourofgo

import "fmt"

type Vertexqw struct {
	X int
	Y int
}

func pointers_to_structs() {
	v := Vertexqw{X: 1, Y: 2}

	p := &v             // p — указатель на v
	p.X = 1_000_000_000 // можно так писать для читаемости

	fmt.Println(v) // {1000000000 2}
	fmt.Println(p) // &{1000000000 2}  ← показывает, что это указатель
}
