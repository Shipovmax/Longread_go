package atourofgo

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl рекурсивно обходит страницы с учётом глубины
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	// TODO: параллельный crawl + защита от дубликатов
	crawl(url, depth, fetcher, make(map[string]bool), &sync.Mutex{})
}

// Вспомогательная функция с visited и mutex
func crawl(url string, depth int, fetcher Fetcher, visited map[string]bool, mu *sync.Mutex) {
	mu.Lock()
	if visited[url] {
		mu.Unlock()
		return
	}
	visited[url] = true
	mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			crawl(u, depth-1, fetcher, visited, mu)
		}(u)
	}

	wg.Wait() // ждём завершения всех горутин на этом уровне
}

func exercise_web_crawler() {
	Crawl("https://golang.org/", 4, fetcher)
}

// ==================== fakeFetcher (оставляем как есть) ====================

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
