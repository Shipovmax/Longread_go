package atourofgo

import "fmt"

// Обычная функция, которая принимает указатель
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func indirection() {
	v := Vertex{3, 4}

	v.Scale(2)        // вызываем метод
	ScaleFunc(&v, 10) // вызываем обычную функцию

	p := &Vertex{4, 3} // p — уже указатель

	p.Scale(3)      // метод
	ScaleFunc(p, 8) // функция

	fmt.Println(v, p)
}
