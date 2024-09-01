package main

import (
	"fmt"
)

func bestSumFactory() func(int, []int) []int {
	var memo = map[int][]int{}

	var bestSum func(int, []int) []int
	bestSum = func(s int, n []int) (h []int) {
		if s == 0 {
			return []int{}
		}

		if s < 0 {
			return nil
		}

		if history, ok := memo[s]; ok {
			return history
		}

		memo[s] = nil

		for _, i := range n {
			if history := bestSum(s-i, n); history != nil {
				if len(h) == 0 || len(history)+1 < len(h) {
					h = append(history, i)
					memo[s] = append(history, i)
				}
			}

		}

		return
	}

	return bestSum
}

func main() {
	var sum int
	var nSlice []int
	fmt.Scanln(&sum)
	for {
		var n int
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}

		nSlice = append(nSlice, n)
	}
	fmt.Println(bestSumFactory()(sum, nSlice))
}
