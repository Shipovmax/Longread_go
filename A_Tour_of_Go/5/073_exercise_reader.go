package atourofgo

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read реализует интерфейс io.Reader
// Возвращает только символы 'A'
func (MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func exercise_reader() {
	reader.Validate(MyReader{})
}
