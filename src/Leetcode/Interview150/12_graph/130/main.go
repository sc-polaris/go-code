package main

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'A'
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}
	// 遍历四周
	for i := 0; i < n; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for j := 1; j < m-1; j++ {
		dfs(0, j)
		dfs(n-1, j)
	}
	for i, row := range board {
		for j, x := range row {
			if x == 'A' {
				board[i][j] = 'O'
			} else if x == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
