// Site statistics with extra convinience methods, like show X most popular words
package main

import (
	"sort"
)

type Statistics struct {
	Words map[string]int
}

// N most popular words with M minimal len
func (statistics *Statistics) MostPopular(minWordLen int, limit int) []struct {
	string
	int
} {
	keys := make([]string, 0, len(statistics.Words))

	for key := range statistics.Words {
		keys = append(keys, key)
	}

	// Sorting only keys
	sort.SliceStable(keys, func(i, j int) bool {
		return statistics.Words[keys[i]] > statistics.Words[keys[j]]
	})

	mostPopular := make([]struct {
		string
		int
	}, 0)
	wordsFound := 0

	for _, k := range keys {
		if len(k) >= minWordLen {
			wordsFound++

			mostPopular = append(mostPopular, struct {
				string
				int
			}{k, statistics.Words[k]})
		}

		if wordsFound == limit {
			return mostPopular
		}
	}

	return mostPopular
}
