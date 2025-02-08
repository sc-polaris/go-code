package main

/*
	给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m - 1][n - 1]）。机器人每次只能向下或者向右移动一步。

	网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。

	返回机器人能够到达右下角的不同路径数量。

	测试用例保证答案小于等于 2 * 10^9。
*/

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
		if *p == -1 { // 没有计算过
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
				f[j+1] += f[j]
			} else {
				f[j+1] = 0
			}
		}
	}
	return f[n]
}
