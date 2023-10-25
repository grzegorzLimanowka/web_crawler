package main

import "time"

func main() {
	go Crawl("https://golang.org/", 3, fetcher)

	// TODO: Sync crawling
	time.Sleep(8 * time.Second)
}
