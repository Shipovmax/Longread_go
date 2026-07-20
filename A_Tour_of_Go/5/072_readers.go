package atourofgo

import (
	"fmt"
	"io"
	"strings"
)

func readers() {
	r := strings.NewReader("Hello, Reader!") // создаёт Reader из строки
	b := make([]byte, 8)                     // буфер на 8 байт для чтения

	for {
		n, err := r.Read(b) // читаем данные в буфер, n — сколько байт прочитано
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF { // проверяем конец данных
			break // выходим из цикла
		}
	}
}
