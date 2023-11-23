package dp

import "math"

const INF = math.MaxInt32

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n, cur := len(grid), len(grid[0]), 0
	dp := [2][]int{
		make([]int, n), make([]int, n),
	}
	copy(dp[0], grid[0])
	for i := 1; i < m; i++ {
		next := 1 - cur
		for j := 0; j < n; j++ {
			dp[next][j] = INF
			for k := 0; k < n; k++ {
				dp[next][j] = min(dp[next][j], dp[cur][k]+moveCost[grid[i-1][k]][j]+grid[i][j])
			}
		}
		cur = next
	}
	res := INF
	for j := 0; j < n; j++ {
		res = min(res, dp[cur][j])
	}

	return res
}
