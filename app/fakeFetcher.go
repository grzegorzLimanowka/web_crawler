package main

import (
	"fmt"
	"log"
)

// TODO: Write Real fetcher, which will scan site and look for valid links for other subsites
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	title      string
	statistics Statistics
	urls       []string // fake hardcoded URLs // TODO: Impl fetching URL from `Base` site.
	// Or even better: Allow to choose strategy for searching - whether to look in hardcoded urls, only from found of both combined
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		statistics, _, err := fetchWords(url)

		if err != nil {
			log.Printf("|error| error fetching words: %v", err)
		}

		mostPopular := statistics.MostPopular(2, 10)
		log.Println("|info| Most popular words: ", mostPopular)

		return res.title, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"Golang site",
		Statistics{},
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Golang packages ",
		Statistics{},
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"formatting section :)",
		Statistics{},
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"op systems",
		Statistics{},
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// TODO: Impl Fetcher, which will be scanning for URLs and searching for new links.
// Add configuration whether it should stick to 'base' url, or go for different domains

// type StatisticsFetcher map[string]*FetchStatistics

// type FetchStatistics struct {
// 	statistics Statistics
// 	urls       []string
// }

// func (f *StatisticsFetcher) Fetch(url string) (string, []string, error) {
// 	statistics, err := fetchWords(url)
// }
