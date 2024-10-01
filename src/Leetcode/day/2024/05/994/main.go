package main

/*
	在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：

	· 值 0 代表空单元格；
	· 值 1 代表新鲜橘子；
	· 值 2 代表腐烂的橘子。
	每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

	返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
*/

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fresh := 0
	var q []pair
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				fresh++
			} else if x == 2 {
				q = append(q, pair{i, j}) // 一开始就腐烂的橘子
			}
		}
	}

	ans := -1
	for len(q) > 0 {
		ans++ // 经过一分钟
		tmp := q
		q = []pair{}
		for _, p := range tmp { // 已经腐烂的橘子
			for _, d := range directions {
				i, j := p.x+d.x, p.y+d.y
				if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 {
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
	return max(ans, 0)
}
