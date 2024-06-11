package main

/*
	给你一个大小为 m x n 的矩阵 board 表示甲板，其中，每个单元格可以是一艘战舰 'X' 或者是一个空位 '.' ，返回在甲板 board 上放置的 战舰 的数量。

	战舰 只能水平或者垂直放置在 board 上。换句话说，战舰只能按 1 x k（1 行，k 列）或 k x 1（k 行，1 列）的形状建造，其中 k 可以是任意大小。两艘
	战舰之间至少有一个水平或垂直的空位分隔 （即没有相邻的战舰）。
*/

/*
	只需要统计战舰头部的个数即可。
	具体来说，如果位于 (i,j) 的格子是战舰的头部，那么左边和上边的相邻格子不能是 X，即：
	· 如果 j > 0，那么 (i,j-1) 不能是 X。
	· 如果 i > 0，那么 (i-1,j) 不能是 X。

*/

func countBattleships(board [][]byte) (ans int) {
	for i, row := range board {
		for j, c := range row {
			if c == 'X' &&
				(j == 0 || board[i][j-1] != 'X') &&
				(i == 0 || board[i-1][j] != 'X') {
				ans++
			}
		}
	}
	return
}
