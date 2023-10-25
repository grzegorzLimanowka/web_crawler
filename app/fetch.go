package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func fetchHTML(URL string) (io.ReadCloser, error) {
	res, err := http.Get(URL)

	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %v", err)
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status code was: %d", res.StatusCode)
	}

	// check content type
	ctype := res.Header.Get("Content-Type")
	if !strings.HasPrefix(ctype, "text/html") {
		return nil, fmt.Errorf("response content type was %s not text/html", ctype)
	}

	return res.Body, nil
}

// TODO: fetchTitle
// This function should extract number of counted words in a `HashMap` format
func extractWords(body io.ReadCloser) (Statistics, []string, error) {

	tokenizer := html.NewTokenizer(body)

	statistics := Statistics{
		Words: make(map[string]int),
	}

	urls := make([]string, 0)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			{
				err := tokenizer.Err()

				if err == io.EOF {
					return statistics, nil, nil
				}

				return statistics, urls, fmt.Errorf("error tokenizing HTML: %v", tokenizer.Err())
			}

		case html.TextToken:
			{
				section := tokenizer.Token().Data

				// TODO: Consider trimming values, remove some unicode garbage?
				for _, val := range strings.Split(section, " ") {
					statistics.Words[val] += 1
				}
			}
		}
		// TODO: Add extracting URLs, preferably only around the 'base' website URL, so we don't go to unrelated sites
		// case html.SelfClosingTagToken:
		// 	{
		//      // TODO..
		// 	}
		// default:
		// 	{
		// 		// fmt.Println("Type", tokenizer.Token().Type)
		// 		// fmt.Println("Attr", tokenizer.Token().Attr)
		// 		// fmt.Println("Data", tokenizer.Token().String())
		// 	}
		// }

	}
}

// TODO: Make it struct implementing interface
func fetchWords(URL string) (Statistics, []string, error) {
	site, err := fetchHTML(URL)

	if err != nil {
		return Statistics{}, nil, err
	}

	defer site.Close()

	statistics, urls, err := extractWords(site)

	if err != nil {
		return statistics, urls, err
	}

	return statistics, urls, nil
}
