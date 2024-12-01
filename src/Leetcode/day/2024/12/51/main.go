package main

import "strings"

/*
	按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

	n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

	给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

	每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

/*
	问：如何 O(1) 判断两个皇后互相攻击？
	答：由于我们保证了每行每列恰好放一个皇后，所以只需检查斜方向。对于 ↗ 方向的格子，行号加列号是不变的。对于 ↖ 方向的格子，行号减列号是不变的。
	   如果两个皇后，行号加列号相同，或者行号减列号相同，那么这两个皇后互相攻击。
	问：如何 O(1) 判断当前位置被之前放置的某个皇后攻击到？
	答：额外用两个数组 diag1 和 diag2 分别标记之前放置的皇后的行号加列号，以及行号减列号。如果当前位置的行号加列号在 diag1 中（标记为 true），
	   或者当前位置的行号减列号在 diag2 中（标记为 true），那么当前位置被之前放置的皇后攻击到，不能放皇后。
*/

func solveNQueens(n int) (ans [][]string) {
	queens := make([]int, n) // 皇后放在 (r,queens[r])
	col := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)
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
		// 在 (r,c) 放皇后
		for c, ok := range col {
			rc := r - c + n - 1
			if !ok && !diag1[r+c] && !diag2[rc] { // 判断能否放皇后
				queens[r] = c                                    // 直接覆盖，无需恢复现场
				col[c], diag1[r+c], diag2[rc] = true, true, true // 皇后占用了 c 列和两条斜线
				dfs(r + 1)
				col[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场
			}
		}
	}
	dfs(0)
	return
}
