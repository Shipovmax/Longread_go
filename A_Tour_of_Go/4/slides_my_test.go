package atourofgo

import "fmt"

func test_my_slides() {
	fmt.Println("=== Работа со слайсами в Go ===\n")

	// 1. Создание слайса через литерал (самый частый способ)
	names := []string{"Макс", "Аня", "Дима", "Катя"}
	fmt.Println("names:", names)

	// 2. Пустой слайс
	var empty []int
	fmt.Println("empty slice:", empty, "len:", len(empty))

	// 3. С помощью make (с предвыделенной памятью)
	scores := make([]int, 0, 10) // len = 0, cap = 10
	scores = append(scores, 10, 20, 30)
	fmt.Println("scores:", scores, "len:", len(scores), "cap:", cap(scores))

	// 4. Слайс структур
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Макс", 25},
		{"Аня", 22},
		{"Игорь", 30},
	}
	fmt.Println("people:", people)

	// 5. Основные операции
	fmt.Println("\n--- Операции ---")
	numbers := []int{1, 2, 3, 4, 5}

	// Добавление
	numbers = append(numbers, 6, 7)
	fmt.Println("После append:", numbers)

	// Срез (slice)
	slice := numbers[2:5]
	fmt.Println("Срез [2:5]:", slice)

	// Изменение через срез влияет на оригинал!
	slice[0] = 999
	fmt.Println("После изменения среза:", numbers)

	// Длина и вместимость
	fmt.Println("len(numbers):", len(numbers))
	fmt.Println("cap(numbers):", cap(numbers))
}
