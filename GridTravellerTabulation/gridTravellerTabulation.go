package main

import (
	"fmt"
)

type gridTraveller struct {
	tabulation [][]int
}

func solveGTTab(row int, col int) int {
	gt := gridTraveller{make([][]int, row+1)}
	for i := range row + 1 {
		gt.tabulation[i] = make([]int, col+1)
	}

	gt.tabulation[1][1] = 1
	for i := range row + 1 {
		for j := range col + 1 {
			if i < row {
				gt.tabulation[i+1][j] += gt.tabulation[i][j]
			}
			if j < col {
				gt.tabulation[i][j+1] += gt.tabulation[i][j]
			}
		}
	}

	return gt.tabulation[row][col]
}

func main() {
	for {
		var i, j int
		fmt.Scan(&i, &j)
		fmt.Println(solveGTTab(i, j))
	}
}
