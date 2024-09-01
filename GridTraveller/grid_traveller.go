package main

import "fmt"

type grid struct {
	max int
	min int
}

func gridTravellerFactory() func(int, int) int {
	var memo = map[grid]int{}

	var gridTraveller func(int, int) int
	gridTraveller = func(w int, h int) int {
		if w == 1 || h == 1 {
			return 1
		}

		g := grid{max(w, h), min(w, h)}

		if r, ok := memo[g]; ok {
			return r
		}

		r := gridTraveller(w-1, h) + gridTraveller(w, h-1)
		memo[g] = r
		return r
	}

	return gridTraveller
}

func main() {
	for {
		var i, j int
		fmt.Scan(&i, &j)
		fmt.Println(gridTravellerFactory()(i, j))
	}
}
