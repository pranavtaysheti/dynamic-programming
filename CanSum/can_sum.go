package main

import (
	"fmt"
)

func canSumFactory() func(s int, n []int) bool {
	var memo = map[int]bool{}

	var canSum func(int, []int) bool
	canSum = func(s int, n []int) bool {
		if s == 0 {
			return true
		}

		if s < 0 {
			return false
		}

		if _, ok := memo[s]; ok {
			return false
		}

		for _, i := range n {
			if canSum(s-i, n) {
				return true
			}
		}

		memo[s] = false
		return false
	}

	return canSum

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
	fmt.Println(canSumFactory()(sum, nSlice))
}
