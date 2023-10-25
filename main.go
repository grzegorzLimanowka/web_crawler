package main

import (
	"fmt"
	"log"
)

func main() {
	statistics, err := fetchWords("https://drstearns.github.io/tutorials/tokenizing/")

	if err != nil {
		log.Fatalf("error fetching words: %v\n", err)
	}

	mostPopular := statistics.MostPopular(2, 10)
	fmt.Println(mostPopular)
}
