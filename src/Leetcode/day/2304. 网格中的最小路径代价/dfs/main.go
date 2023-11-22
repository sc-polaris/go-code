package dfs

import "math"

const INF = math.MaxInt32

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m) // 记录第0行任意单元格到第i行j列的最小路径代价
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i, j int) int {
		if i == 0 {
			return grid[i][j]
		}

		if memo[i][j] >= 0 {
			return memo[i][j]
		}
		res := INF
		for k := 0; k < n; k++ { // 遍历列
			res = min(res, dfs(i-1, k)+moveCost[grid[i-1][k]][j]+grid[i][j])
		}
		memo[i][j] = res

		return res
	}

	res := INF
	for j := 0; j < n; j++ { // 遍历行
		res = min(res, dfs(m-1, j))
	}

	return res
}
