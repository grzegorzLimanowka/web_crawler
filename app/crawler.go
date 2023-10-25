package main

// TODO: Find formatted logger
import (
	"log"
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
	// Fetch returns the body of URL and a slice of URLs found on that page
	Fetch(url string) (title string, urls []string, err error)
}

func CrawlInternal(visited *VisitedSites, url string, depth int, fetcher Fetcher) {
	if depth < 0 {
		return
	}

	if visited.Get(url) {
		log.Printf("|Trace| - Ingnoring url %s, which has been visited before ...\n", url)
		return
	} else {
		visited.Add(url)
	}

	title, urls, err := fetcher.Fetch(url)

	if err != nil {
		log.Printf("|Warn| - %s", err)
		return
	}

	log.Printf("|Info| - Found %s %q\n", url, title)

	for _, u := range urls {
		go CrawlInternal(visited, u, depth-1, fetcher)
	}
}

// Start crawling for X depth
func Crawl(url string, depth int, fetcher Fetcher) {
	visited := VisitedSites{
		mu:      sync.Mutex{},
		visited: make(map[string]bool),
	}

	CrawlInternal(&visited, url, depth, fetcher)
}
