package atourofgo

import (
	"fmt"
	"math"
)

func poww(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func if_and_else() {
	fmt.Println(
		poww(3, 2, 10),
		poww(3, 3, 20),
	)
}
