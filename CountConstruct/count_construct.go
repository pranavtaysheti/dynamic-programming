package main

import (
	"fmt"
)

type countConstruct struct {
	memo map[string]int
}

func solveCountConstruct(s string, wb []string) int {
	countConstruct := countConstruct{
		memo: map[string]int{},
	}

	return countConstruct.solve(s, wb)
}

func (c countConstruct) solve(s string, wb []string) int {
	if s == "" {
		return 1
	}

	if ways, ok := c.memo[s]; ok {
		return ways
	}

	total_ways := 0
	for _, word := range wb {
		if len(word) > len(s) {
			continue
		}

		if s[:len(word)] == word {
			total_ways += c.solve(s[len(word):], wb)
		}
	}

	c.memo[s] = total_ways
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
	fmt.Println(solveCountConstruct(sentence, words))
}
