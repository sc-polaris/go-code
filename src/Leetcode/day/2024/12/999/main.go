package main

/*
	给定一个 8 x 8 的棋盘，只有一个 白色的车，用字符 'R' 表示。棋盘上还可能存在白色的象 'B' 以及黑色的卒 'p'。空方块用字符 '.' 表示。

	车可以按水平或竖直方向（上，下，左，右）移动任意个方格直到它遇到另一个棋子或棋盘的边界。如果它能够在一次移动中移动到棋子的方格，则能够 吃掉 棋子。

	注意：车不能穿过其它棋子，比如象和卒。这意味着如果有其它棋子挡住了路径，车就不能够吃掉棋子。

	返回白车 攻击 范围内 兵的数量。
*/

type dir struct{ x, y int }

var dirs = []dir{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func numRookCaptures(board [][]byte) (ans int) {
	const size = 8
	var x0, y0 int
	for i, row := range board {
		for j, c := range row {
			if c == 'R' {
				x0, y0 = i, j
			}
		}
	}
	for _, d := range dirs {
		x, y := x0+d.x, y0+d.y
		for 0 <= x && x < size && 0 <= y && y < size && board[x][y] == '.' {
			x += d.x
			y += d.y
		}
		if 0 <= x && x < size && 0 <= y && y < size && board[x][y] == 'p' {
			ans++
		}
	}
	return
}
