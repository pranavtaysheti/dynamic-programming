package main

import (
	"fmt"
)

type howSum struct {
	tabulation [][]int
}

func solveHowSum(t int, nb []int) []int {
	hs := howSum{
		tabulation: make([][]int, t+1),
	}

	hs.tabulation[0] = []int{}

	for i := range t + 1 {
		if hs.tabulation[t] != nil {
			return hs.tabulation[t]
		}

		if hs.tabulation[i] == nil {
			continue
		}

		for _, n := range nb {
			if i+n <= t {
				hs.tabulation[i+n] = append(hs.tabulation[i], n)
			}
		}
	}

	return nil
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

	res := solveHowSum(sum, nSlice)

	if res == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(res)
}
