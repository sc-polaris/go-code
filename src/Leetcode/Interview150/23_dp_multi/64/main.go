package main

import "math"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if i == 0 && j == 0 {
			return grid[i][j]
		}
		p := &memo[i][j]
		if *p == -1 {
			*p = min(dfs(i, j-1), dfs(i-1, j)) + grid[i][j]
		}
		return *p
	}
	return dfs(m-1, n-1)
}

func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for j := 2; j <= n; j++ {
		f[0][j] = math.MaxInt
	}
	for i, row := range grid {
		f[i+1][0] = math.MaxInt
		for j, x := range row {
			f[i+1][j+1] = min(f[i+1][j], f[i][j+1]) + x
		}
	}
	return f[m][n]
}

func minPathSum3(grid [][]int) int {
	n := len(grid[0])
	f := make([]int, n+1)
	for j := range f {
		f[j] = math.MaxInt
	}
	f[1] = 0
	for _, row := range grid {
		for j, x := range row {
			f[j+1] = min(f[j], f[j+1]) + x
		}
	}
	return f[n]
}
