package main

import (
	"fmt"
)

type allConstruct struct {
	memo map[string][][]string
}

func solveAllConstruct(t string, wb []string) [][]string {
	allConstruct := allConstruct{
		memo: map[string][][]string{},
	}

	return allConstruct.solve(t, wb)
}

func (c allConstruct) solve(t string, wb []string) [][]string {
	if t == "" {
		return [][]string{{}}
	}

	if ways, ok := c.memo[t]; ok {
		return ways
	}

	total_ways := [][]string{}
	for _, word := range wb {
		if len(word) > len(t) {
			continue
		}

		if t[:len(word)] == word {
			for _, way := range c.solve(t[len(word):], wb) {
				wordSlice := []string{word}
				total_ways = append(total_ways, append(wordSlice, way...))
			}
		}
	}

	c.memo[t] = total_ways
	return total_ways
}

func main() {
	var sentence string
	var words []string
	fmt.Scanln(&sentence)
	for {
		var word string
		_, err := fmt.Scanf("%s", &word)
		if err != nil {
			break
		}

		words = append(words, word)
	}
	fmt.Println(solveAllConstruct(sentence, words))
}
