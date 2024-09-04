package main

func gameOfLife(board [][]int) {
	n, m := len(board), len(board[0])
	dx := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	dy := []int{0, 0, -1, 1, -1, 1, 1, -1}

	for i, row := range board {
		for j := range row {
			live := 0
			for k := range dx {
				x, y := i+dx[k], j+dy[k]
				if x < 0 || x >= n || y < 0 || y >= m {
					continue
				}
				// // 这里的0,1是题目里合理的值，然后如果0要变1，我们用中间值-1记录，如果1要变0，我们用中间值2来记录。
				if board[x][y] == 1 || board[x][y] == 2 {
					live++
				}
			}
			if board[i][j] == 0 && live == 3 { // 死亡变活
				board[i][j] = -1
			} else if board[i][j] == 1 && (live < 2 || live > 3) { // 活变死亡
				board[i][j] = 2
			}
			// 仍然存活省略
			//else if board[i][j] == 1 && (live == 2 || live == 3) {
			//	board[i][j] = 1
			//}
		}
	}

	for i, row := range board {
		for j := range row {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}

// golang 特性
func gameOfLife2(board [][]int) {
	dx := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	dy := []int{0, 0, -1, 1, -1, 1, 1, -1}
	getRes := func(i, j int) int {
		live := 0
		for k := range dx {
			x := i + dx[k]
			y := j + dy[k]
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[i]) {
				live += board[x][y]
			}
		}

		if live < 2 || live > 3 {
			return 0
		} else if live == 3 && board[i][j] == 0 {
			return 1
		}

		return board[i][j]
	}

	for i, row := range board {
		for j := range row {
			res := getRes(i, j)
			defer func(i, j, res int) { // defer 延迟执行
				board[i][j] = res
			}(i, j, res)
		}
	}
}
