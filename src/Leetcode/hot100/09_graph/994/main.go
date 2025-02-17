package main

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	fresh := 0
	var q []pair
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				fresh++
			} else if x == 2 {
				q = append(q, pair{i, j})
			}
		}
	}

	ans := 0
	for fresh > 0 && len(q) > 0 {
		ans++
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, d := range dirs {
				i, j := p.x+d.x, p.y+d.y
				if 0 <= i && i < n && 0 <= j && j < m && grid[i][j] == 1 {
					fresh--
					grid[i][j] = 2
					q = append(q, pair{i, j})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return ans
}
