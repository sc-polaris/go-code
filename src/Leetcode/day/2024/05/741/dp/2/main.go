package main

/*
	在计算 f[t] 时，只会用到 f[t−1]，不会用到比 t−1 更早的状态
	因此可以像 0-1 背包一样，去掉第一个纬度，倒序枚举 j 和 k。

	状态转移方程改为
		f[j+1][k+1] = max(f[j+1][k+1],f[j+1][k],f[j][k+1],f[j][k]) + val
	初始值 f[j][k] = 负无穷，f[1][1] = grid[0][0]。
	答案为 max(f[n][n],0)。
*/

import "math"

func cherryPickup(grid [][]int) int {
	n := len(grid)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = math.MinInt
		}
	}
	f[1][1] = grid[0][0]
	for t := 1; t < n*2-1; t++ {
		for j := min(t, n-1); j >= max(t-n+1, 0); j-- {
			for k := min(t, n-1); k >= j; k-- {
				if grid[t-j][j] < 0 || grid[t-k][k] < 0 {
					f[j+1][k+1] = math.MinInt
					continue
				}
				f[j+1][k+1] = max(max(f[j+1][k+1], f[j+1][k]), max(f[j][k+1], f[j][k])) + grid[t-j][j]
				if k != j {
					f[j+1][k+1] += grid[t-k][k]
				}
			}
		}
	}
	return max(f[n][n], 0)
}
