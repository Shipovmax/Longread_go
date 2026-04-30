package atourofgo

import "fmt"

func Sqrt(x float64) float64 {
	z := 1.0                  // начальное приближение
	for i := 0; i < 10; i++ { // обычно 10 итераций достаточно
		z = z - (z*z-x)/(2*z) // формула Ньютона
	}
	return z
}

func exercise_loops_and_functions() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(16))
}
