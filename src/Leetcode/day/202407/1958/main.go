package main

/*
	给你一个下标从 0 开始的 8 x 8 网格 board ，其中 board[r][c] 表示游戏棋盘上的格子 (r, c) 。棋盘上空格用 '.' 表示，白色格子用 'W' 表示，黑色格子用 'B' 表示。

	游戏中每次操作步骤为：选择一个空格子，将它变成你正在执行的颜色（要么白色，要么黑色）。但是，合法 操作必须满足：涂色后这个格子是 好线段的一个端点 （好线段可以是水平的，
	竖直的或者是对角线）。

	好线段 指的是一个包含 三个或者更多格子（包含端点格子）的线段，线段两个端点格子为 同一种颜色 ，且中间剩余格子的颜色都为 另一种颜色 （线段上不能有任何空格子）。

	给你两个整数 rMove 和 cMove 以及一个字符 color ，表示你正在执行操作的颜色（白或者黑），如果将格子 (rMove, cMove) 变成颜色 color 后，是一个 合法 操作，那么返回
	true ，如果不是合法操作返回 false 。
*/

/*
	思路：
	枚举八个方向：
	1. 该方向的下一个格子不能出界，颜色必须和 color 相反。
	2. 从下下一个格子开始，在遇到颜色等于 color 的格子之前，不能出界，不能遇到空格子。
	3. 如果满足前两个要求，则在遇到颜色等于 color 的格子时返回 true。

	代码技巧：
	用 color ^ 'B' ^ 'W' 表示「和 color 相反的颜色」：
	· 当 color 是 B 时，会计算出 W。
	· 当 color 是 W 时，会计算出 B。

	注：该技巧仅限于本题简化代码，由于可读性的原因，工程上可以用 board[x][y] == '.' || board[x][y] == color 代替 board[x][y] != color ^ 'B' ^ 'W'。

*/

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	dirs := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	m, n := len(board), len(board[0])
	for _, dir := range dirs {
		x, y := rMove+dir.x, cMove+dir.y
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] == '.' || board[x][y] == color {
			continue
		}
		for {
			x += dir.x
			y += dir.y
			if x < 0 || x >= m || y < 0 || y >= n || board[x][y] == '.' {
				break
			}
			if board[x][y] == color {
				return true
			}
		}
	}
	return false
}
