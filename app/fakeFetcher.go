package main

import "fmt"

// TODO: Write Real fetcher, which will scan site and look for valid links for other subsites
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body Statistics // empty statistics
	urls []string   // fake hardcoded URLs
}

func (f fakeFetcher) Fetch(url string) (Statistics, []string, error) {
	if res, ok := f[url]; ok {
		fmt.Println("BODY", res.body)

		return res.body, res.urls, nil
	}

	return Statistics{}, nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		Statistics{},
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		Statistics{},
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		Statistics{},
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		Statistics{},
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
