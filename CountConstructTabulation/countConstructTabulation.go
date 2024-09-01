package main

import (
	"fmt"
)

type tabulation []int

func solveCountConstruct(t string, wb []string) int {
	cc := tabulation(make([]int, len(t)+1))
	cc[0] = 1

	for i := range len(t) + 1 {
		if cc[i] == 0 {
			continue
		}

		for _, w := range wb {
			if i+len(w) > len(t) {
				continue
			}

			if w == t[i:i+len(w)] {
				cc[i+len(w)] += cc[i]
			}
		}
	}

	return cc[len(t)]
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
