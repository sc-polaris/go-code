package main

import "math"

/*
	四：1:1 翻译成递推
	f[t][j][k] 定义和 dfs(t,j,k) 是一样的，都表示两人从 (0,0) 出发，都走了 t 步，分别走到 (t-j,j) 和 (t-k,k)，可以的到的樱桃的最大值

	f[t][j][k] = max(f[t-1][j][k],f[t-1][j][k-1],f[t-1][j-1][k],f[t-1][j-1][k-1]) + val
	其中：
	val        = grid[t-j][j] + grid[t-k][k]   j != k
			   = grid[t-j][j]				   j == k

	但是，这种定义方式没有状态能表示递归边界，即 j = -1, k = -1 这种出界的情况。
	解决方法：在每个 f[i] 的最左边和最上边各插入一排状态，那么其余状态全部向右和向下偏移一位，把 f[t][j][k] 改为 f[t][j+1][k+1]。
	修改后 f[t][j+1][k+1] 表示两人从 (0,0) 出发，都走了 t 步，分别走到 (t-j,j) 和 (t-k,k)，可以得到的樱桃个数的最大值。此时 f[t][0][·]
	和 f[t][·][0] 就对应出界的情况了。

	修改后的递推式：
	f[t][j+1][k+1] = max(f[t-1][j+1][k+1],f[t-1][j+1][k],f[t-1][j][k+1],f[t-1][j][k]) + val

	注意我们只在 f 数组上插入了状态，这只会影响 f 的下标，val 的计算方式不变。
	初始值 f[t][j][k] = 负无穷，f[0][1][1] = grid[0][0]
	答案为 max(f[2n-2][n][n],0)。

	循环范围：
	代码实现时，我们还需要讨论清楚 j 和 k 的范围。
	由于 i + j = t 且 0 <= i <= n-1 且 0 <= j <= n-1，联立得
		max(t-n+1,0) <= j <= min(t,n-1）
	对于 k 也同理。
*/

func cherryPickup(grid [][]int) int {
	n := len(grid)
	f := make([][][]int, n*2-1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, n+1)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt
			}
		}
	}
	f[0][1][1] = grid[0][0]
	for t := 1; t < n*2-1; t++ {
		for j := max(t-n+1, 0); j <= min(t, n-1); j++ {
			if grid[t-j][j] < 0 {
				continue
			}
			for k := j; k <= min(t, n-1); k++ {
				if grid[t-k][k] < 0 {
					continue
				}
				f[t][j+1][k+1] = max(f[t-1][j+1][k+1], f[t-1][j+1][k], f[t-1][j][k+1], f[t-1][j][k]) + grid[t-j][j]
				if k != j {
					f[t][j+1][k+1] += grid[t-k][k]
				}
			}
		}
	}
	return max(f[n*2-2][n][n], 0)
}
