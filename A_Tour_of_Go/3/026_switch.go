package atourofgo

import (
	"fmt"
	"runtime"
)

func switch_test() {
	// runtime.GOOS — встроенная константа, содержащая название ОС
	currentOS := runtime.GOOS
	fmt.Printf("Go runs on %s.\n", currentOS)

	// Switch без выражения (сравнивает напрямую с case)
	switch currentOS {

	case "darwin": // macOS / iOS
		fmt.Println("Привет, любитель macOS! 🍎")

	case "linux":
		fmt.Println("Привет, Linux-воин! 🐧")

	case "windows":
		fmt.Println("Привет, пользователь Windows! 🪟")

	default:
		// Сюда попадут: freebsd, android, openbsd и другие
		fmt.Printf("Привет, экзотическая ОС: %s!\n", currentOS)
	}
}
