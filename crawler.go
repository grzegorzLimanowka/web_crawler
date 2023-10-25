package main

import (
	"fmt"
	"sync"
)

type VisitedSites struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (v *VisitedSites) Add(visited string) {
	v.mu.Lock()
	v.visited[visited] = true
	v.mu.Unlock()
}

func (v *VisitedSites) Get(visited string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.visited[visited]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

func CrawlInternal(visited *VisitedSites, url string, depth int, fetcher Fetcher) {
	if depth < 0 {
		return
	}

	if visited.Get(url) {
		fmt.Printf("Ingnoring url %s, which has been visited before ...\n", url)
		return
	} else {
		visited.Add(url)
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found %s %q\n", url, body)

	for _, u := range urls {
		go CrawlInternal(visited, u, depth-1, fetcher)
	}
}

func Crawl(url string, depth int, fetcher Fetcher) {
	visited := VisitedSites{
		mu:      sync.Mutex{},
		visited: make(map[string]bool),
	}

	CrawlInternal(&visited, url, depth, fetcher)
}

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

// Real fetcher should fetch from real URL i guess

// func main() {
// 	go Crawl("https://golang.org/", 4, fetcher)

// 	time.Sleep(8 * time.Second)
// }

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
