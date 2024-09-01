package main

import (
	"fmt"
)

type tabulation [][][]string

func solveAllConstruct(t string, wb []string) [][]string {
	cc := tabulation(make([][][]string, len(t)+1))
	cc[0] = [][]string{{}}

	for i := range len(t) + 1 {
		if cc[i] == nil {
			continue
		}

		for _, w := range wb {
			if i+len(w) > len(t) {
				continue
			}

			if w == t[i:i+len(w)] {
				for _, s := range cc[i] {
					cc[i+len(w)] = append(cc[i+len(w)], append(s, w))
				}
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
	fmt.Println(solveAllConstruct(sentence, words))
}
