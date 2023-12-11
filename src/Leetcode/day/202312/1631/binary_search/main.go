package main

import (
	"sort"
)

func minimumEffortPath(heights [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, m := len(heights), len(heights[0])
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		st := make([][]bool, n)
		for i := range st {
			st[i] = make([]bool, m)
		}
		st[0][0] = true
		q := []pair{{0, 0}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == n-1 && p.y == m-1 {
				return true
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < m && !st[x][y] && abs(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
					st[x][y] = true
					q = append(q, pair{x, y})
				}
			}
		}
		return false
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
