package main

import (
	"fmt"
)

type canConstruct struct {
	memo map[string]bool
}

func solveCanConstruct(s string, wb []string) bool {
	canConstruct := canConstruct{
		memo: map[string]bool{},
	}

	return canConstruct.solve(s, wb)
}

func (c canConstruct) solve(s string, wb []string) bool {
	if s == "" {
		return true
	}

	if _, ok := c.memo[s]; ok {
		return false
	}

	for _, w := range wb {
		if len(w) > len(s) {
			continue
		}

		if s[:len(w)] == w && c.solve(s[len(w):], wb) {
			return true
		}
	}

	c.memo[s] = false
	return false

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
	fmt.Println(solveCanConstruct(sentence, words))
}
