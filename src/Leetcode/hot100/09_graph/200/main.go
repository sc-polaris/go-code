package main

func numIslands(grid [][]byte) (ans int) {
	n, m := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i, row := range grid {
		for j, x := range row {
			if x == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}
