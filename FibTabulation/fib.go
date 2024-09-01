package main

import (
	"fmt"
)

type tabulation []int

func solveFibTabulation(t int) int {
	ft := tabulation(make([]int, t+1))
	ft[0] = 0
	ft[1] = 1

	for i := range t - 1 {
		c := i + 2
		ft[c] = ft[c-1] + ft[c-2]
	}

	return ft[t]
}

func main() {
	for {
		var i int
		fmt.Scan(&i)
		fmt.Println(solveFibTabulation(i))
	}
}
