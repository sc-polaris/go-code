package main

/*
	在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。行和列是 从 0 开始 的，所以左上单元格是 (0,0) ，
	右下单元格是 (n - 1, n - 1) 。

	象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。

	每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。

	骑士继续移动，直到它走了 k 步或离开了棋盘。

	返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。
*/

/*
	dfs
	k：还剩下 k 步要走。
	(i,j)：马的位置。
	因此，定义状态为 dfs(k,i,j)，表示马从 (i,j) 出发，走 k 步后仍然在棋盘上的概率。

	枚举马走的八个方向，其中有 1/8 概率走到了 (x,y)，问题变成：
	· 马从 (x,y) 出发，走 k−1 步后仍然在棋盘上的概率，即 dfs(k−1,x,y)。
*/

func knightProbability(n int, k int, row int, column int) float64 {
	var dirs = []struct{ x, y int }{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	memo := make([][][]float64, k+1)
	for i := range memo {
		memo[i] = make([][]float64, n)
		for j := range memo[i] {
			memo[i][j] = make([]float64, n)
		}
	}
	var dfs func(int, int, int) float64
	dfs = func(k, i, j int) float64 {
		if i < 0 || j < 0 || i >= n || j >= n {
			return 0
		}
		if k == 0 {
			return 1
		}
		p := &memo[k][i][j]
		if *p > 0 {
			return *p
		}
		res := 0.0
		for _, d := range dirs {
			res += dfs(k-1, i+d.x, j+d.y)
		}
		res /= 8
		*p = res
		return res
	}
	return dfs(k, row, column)
}

/*
	为避免下标出现负数，把棋盘的边界从 [0,n−1] 调整为 [2,n+1]。
*/

func knightProbability2(n int, k int, row int, column int) float64 {
	var dirs = []struct{ x, y int }{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	f := make([][][]float64, k+1)
	for i := range f {
		f[i] = make([][]float64, n+4)
		for j := range f[i] {
			f[i][j] = make([]float64, n+4)
		}
	}
	for i := 2; i < n+2; i++ {
		for j := 2; j < n+2; j++ {
			f[0][i][j] = 1
		}
	}
	for step := 1; step <= k; step++ {
		for i := 2; i < n+2; i++ {
			for j := 2; j < n+2; j++ {
				for _, d := range dirs {
					f[step][i][j] += f[step-1][i+d.x][j+d.y]
				}
				f[step][i][j] /= 8
			}
		}
	}
	return f[k][row+2][column+2]
}
