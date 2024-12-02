package main

/*
	n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

	给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
*/

func totalNQueens(n int) (ans int) {
	col := make([]bool, n)       // 列
	diag1 := make([]bool, n*2-1) // 左对角线
	diag2 := make([]bool, n*2-1) // 右对角线
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			ans++
			return
		}
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !diag1[r+c] && !diag2[rc] {
				col[c], diag1[r+c], diag2[rc] = true, true, true
				dfs(r + 1)
				col[c], diag1[r+c], diag2[rc] = false, false, false // 回溯
			}
		}
	}
	dfs(0)
	return
}
