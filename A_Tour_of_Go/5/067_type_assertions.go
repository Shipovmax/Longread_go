package atourofgo

import "fmt"

// Speaker — интерфейс, который описывает любой тип, способный "говорить"
type Speaker interface {
	Speak()
}

// Message — структура, которая содержит текст и умеет его выводить
type Message struct {
	Text string
}

// Speak — метод, который выводит содержимое сообщения
func (m *Message) Speak() {
	if m == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(m.Text)
}

// showInfo — выводит информацию о том, что лежит внутри интерфейса
func showInfo(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func type_assertions() {
	var holder interface{} // переменная пустого интерфейса

	showInfo(holder)

	holder = 42
	showInfo(holder)

	holder = "hello"
	showInfo(holder)

	holder = &Message{Text: "World"}
	showInfo(holder)
	holder.(Speaker).Speak() // World
}
