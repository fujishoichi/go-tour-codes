package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		wordCountMap[word]++

		// これもok
		//_, ok := wordCountMap[word]
		//if ok {
		//	wordCountMap[word]++
		//} else {
		//	wordCountMap[word] = 1
		//}
	}
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
