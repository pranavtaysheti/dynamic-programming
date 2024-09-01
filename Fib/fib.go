package main

import (
	"fmt"
)

type memo map[int]int

func solveFib(i int) int {
	f := memo(map[int]int{})

	return f.calculate(i)
}

func (f memo) calculate(i int) int {
	if i == 1 || i == 2 {
		return 1
	}

	if sum, ok := f[i]; ok {
		return sum
	}

	res := f.calculate(i-1) + f.calculate(i-2)
	f[i] = res
	return res
}

func main() {
	for {
		var i int
		fmt.Scan(&i)
		fmt.Println(solveFib(i))
	}
}
