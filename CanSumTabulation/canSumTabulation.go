package main

import (
	"fmt"
)

type canSum struct {
	tabulation []bool
}

func solveCanSum(t int, nb []int) bool {
	cs := canSum{make([]bool, t+1)}
	cs.tabulation[0] = true

	for i := range t + 1 {
		if cs.tabulation[t] {
			return true
		}

		if !cs.tabulation[i] {
			continue
		}

		for _, n := range nb {
			if i+n <= t {
				cs.tabulation[i+n] = true
			}
		}
	}

	return false
}
func main() {
	for {
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

		fmt.Println(solveCanSum(sum, nSlice))
	}
}
