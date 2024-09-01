package main

import (
	"fmt"
)

type howSum struct {
	tabulation [][]int
}

func solveHowSum(t int, nb []int) []int {
	bs := howSum{
		tabulation: make([][]int, t+1),
	}

	bs.tabulation[0] = []int{}

	for i := range t + 1 {
		if bs.tabulation[i] == nil {
			continue
		}

		for _, n := range nb {
			if i+n <= t && (len(bs.tabulation[i+n]) > len(bs.tabulation[i])+1 || bs.tabulation[i+n] == nil) {
				bs.tabulation[i+n] = append(bs.tabulation[i], n)
			}
		}
	}

	return bs.tabulation[t]
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
