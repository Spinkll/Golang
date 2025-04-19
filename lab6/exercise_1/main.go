package main

import (
	"fmt"
	"sync"
)

type Fetcher interface { // Fetcher — інтерфейс для завантаження сторінок.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct { // Crawl — паралельно завантажує сторінки, використовуючи Fetcher.
	mu sync.Mutex
	v  map[string]bool
}

// CheckAndSet перевіряє, чи вже був відвіданий URL.
// Якщо ні, додає його до списку відвіданих та повертає true.
func (s *SafeMap) CheckAndSet(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.v[url] {
		return false // URL вже відвідували, не обробляємо повторно.
	}
	s.v[url] = true
	return true
}

func Crawl(url string, depth int, fetcher Fetcher, visited *SafeMap, wg *sync.WaitGroup) { // Crawl — паралельно завантажує сторінки, використовуючи Fetcher.
	defer wg.Done() // Зменшує лічильник goroutine після завершення.

	if depth <= 0 || !visited.CheckAndSet(url) { // Якщо досягли максимальної глибини або URL вже оброблено — виходимо.
		return
	}

	body, urls, err := fetcher.Fetch(url) // Викликаємо Fetch для отримання сторінки.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body) // Виводимо інформацію про знайдену сторінку.

	wg.Add(len(urls))

	for _, u := range urls { // Запускаємо нові горутини для кожного знайденого URL.
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
}

func main() {
	var wg sync.WaitGroup
	visited := &SafeMap{v: make(map[string]bool)} // Ініціалізація структури для перевірки відвіданих URL.

	wg.Add(1) // Додаємо першу задачу.
	go Crawl("https://golang.org/", 4, fetcher, visited, &wg)

	wg.Wait() // Очікуємо завершення всіх горутин.
}

type fakeFetcher map[string]*fakeResult // fakeFetcher — реалізація Fetcher, що повертає тестові дані.

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) { // Fetch — імітує отримання сторінки
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
