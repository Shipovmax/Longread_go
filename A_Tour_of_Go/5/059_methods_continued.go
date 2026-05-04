package atourofgo

import (
	"fmt"
	"math"
)

// MyFloat — это новый тип на основе float64
type MyFloat float64

// Abs — метод, привязанный к типу MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f) // делаем число положительным
	}
	return float64(f)
}

func methods_continued() {
	f := MyFloat(-math.Sqrt2) // ≈ -1.41421356237
	fmt.Println(f.Abs())      // → 1.41421356237
}
