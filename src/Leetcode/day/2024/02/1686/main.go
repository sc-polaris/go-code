package main

import (
	"cmp"
	"slices"
)

func stoneGameVI(a []int, b []int) int {
	type pair struct{ x, y int }
	pairs := make([]pair, len(a))
	for i, x := range a {
		pairs[i] = pair{x, b[i]}
	}
	slices.SortFunc(pairs, func(p, q pair) int { return q.x + q.y - p.x - p.y })
	diff := 0
	for i, p := range pairs {
		if i&1 == 0 {
			diff += p.x
		} else {
			diff -= p.y
		}
	}
	return cmp.Compare(diff, 0)
}

func main() {

}
