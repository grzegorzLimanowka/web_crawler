package main

import (
	"time"
)

func main() {
	// statistics, err := fetchWords("https://drstearns.github.io/tutorials/tokenizing/")

	// if err != nil {
	// 	log.Fatalf("error fetching words: %v\n", err)
	// }

	// mostPopular := statistics.MostPopular(2, 10)
	// fmt.Println(mostPopular)

	//

	go Crawl("https://golang.org/", 3, fetcher)

	time.Sleep(8 * time.Second)

}
