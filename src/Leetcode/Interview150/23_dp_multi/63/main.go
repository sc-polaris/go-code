package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || obstacleGrid[i][j] == 1 {
			return 0
		}
		if i == 0 && j == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p == -1 {
			*p = dfs(i-1, j) + dfs(i, j-1)
		}
		return *p
	}
	return dfs(m-1, n-1)
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][1] = 1
	for i, row := range obstacleGrid {
		for j, x := range row {
			if x == 0 {
				f[i+1][j+1] = f[i][j+1] + f[i+1][j]
			}
		}
	}
	return f[m][n]
}

func uniquePathsWithObstacles3(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	f := make([]int, n+1)
	f[1] = 1
	for _, row := range obstacleGrid {
		for j, x := range row {
			if x == 0 {
				f[j+1] = f[j] + f[j+1]
			} else {
				f[j+1] = 0
			}
		}
	}
	return f[n]
}
