package main

import (
	"fmt"
)

func howSumFactory() func(int, []int) []int {
	var memo = map[int]bool{}

	var howSum func(int, []int) []int
	howSum = func(s int, n []int) []int {
		if s == 0 {
			return []int{}
		}

		if s < 0 {
			return nil
		}

		if _, ok := memo[s]; ok {
			return nil
		}

		for _, i := range n {
			if history := howSum(s-i, n); history != nil {
				return append(history, i)
			}
		}

		memo[s] = false
		return nil
	}

	return howSum
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
	fmt.Println(howSumFactory()(sum, nSlice))
}
