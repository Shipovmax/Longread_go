package atourofgo

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)    // разбиваем строку на слова
	count := make(map[string]int) // создаём map для подсчёта

	for _, word := range words {
		count[word]++ // увеличиваем счётчик для этого слова
	}

	return count
}

func exercise_maps() {
	wc.Test(WordCount)
}
