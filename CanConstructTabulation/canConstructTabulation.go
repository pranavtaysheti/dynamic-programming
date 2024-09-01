package main

import (
	"fmt"
)

type tabulation []bool

func solveCanConstruct(t string, wb []string) bool {
	cc := tabulation(make([]bool, len(t)+1))
	cc[0] = true

	for i := range len(t) + 1 {
		if cc[len(t)] {
			return true
		}

		if !cc[i] {
			continue
		}

		for _, w := range wb {
			if i+len(w) > len(t) {
				continue
			}

			if w == t[i:i+len(w)] {
				cc[i+len(w)] = true
			}
		}
	}

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
