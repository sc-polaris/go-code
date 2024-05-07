package main

/*
	五、空间优化：滚动数组
	观察上边的状态转移方程，在计算 f[i] 时，只会用到 f[i+1]，不会用到 > i+1 的状态。
	我们可以用两个二位数组滚动计算，用 cur 表示 f[i]，用 pre 表示 f[i+1]，状态转移方式改为：
		cur[j+1][k+1] = val + max { pre[j][k], pre[j][k+1], pre[j][k+2],
									pre[j+1][k], pre[j+1][k+1], pre[j+1][k+2],
									pre[j+2][k], pre[j+2][k+1], pre[j+2][k+2] }

	从 i 枚举到 i-1 之前，交换 cur 和 pre，相当于把 cur 变成对 i-1 而言的 pre。
*/

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	pre := make([][]int, n+2)
	cur := make([][]int, n+2)
	for i := range pre {
		pre[i] = make([]int, n+2)
		cur[i] = make([]int, n+2)
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < min(n, i+1); j++ {
			for k := max(j+1, n-1-i); k < n; k++ {
				cur[j+1][k+1] = max(
					pre[j][k], pre[j][k+1], pre[j][k+2],
					pre[j+1][k], pre[j+1][k+1], pre[j+1][k+2],
					pre[j+2][k], pre[j+2][k+1], pre[j+2][k+2],
				) + grid[i][j] + grid[i][k]
			}
		}
		pre, cur = cur, pre // 下一个 i 的 pre 是 cur
	}
	return pre[1][n]
}
