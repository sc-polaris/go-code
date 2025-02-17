package main

import "strings"

func solveNQueens(n int) (ans [][]string) {
	queens := make([]int, n) // 皇后放在 (r,queens[r])
	col := make([]bool, n)
	d1 := make([]bool, n*2-1)
	d2 := make([]bool, n*2-1)
	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range queens {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board)
			return
		}
		// 在 (r, c) 放皇后
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !d1[r+c] && !d2[rc] {
				queens[r] = c
				col[c], d1[r+c], d2[rc] = true, true, true
				dfs(r + 1)
				col[c], d1[r+c], d2[rc] = false, false, false
			}
		}
	}
	dfs(0)
	return
}
