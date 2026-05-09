package atourofgo

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader // оборачиваем любой Reader
}

// Read — главная функция, которая реализует интерфейс io.Reader
// Читает данные из внутреннего Reader, применяет ROT13 и возвращает результат
func (rot rot13Reader) Read(b []byte) (int, error) {
	// Сначала читаем оригинальные данные во временный буфер
	n, err := rot.r.Read(b)

	// Проходим по всем прочитанным байтам и расшифровываем ROT13
	for i := 0; i < n; i++ {
		c := b[i]

		// Для заглавных букв A-Z
		if c >= 'A' && c <= 'Z' {
			b[i] = 'A' + (c-'A'+13)%26
		} else if c >= 'a' && c <= 'z' { // Для строчных букв a-z
			b[i] = 'a' + (c-'a'+13)%26
		}
		// Остальные символы (пробелы, знаки и т.д.) оставляем как есть
	}

	return n, err
}

func exercise_rot_reader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // зашифрованная строка
	r := rot13Reader{s}                             // оборачиваем в наш rot13Reader

	// io.Copy читает из r до конца и пишет в os.Stdout
	io.Copy(os.Stdout, &r)
}
