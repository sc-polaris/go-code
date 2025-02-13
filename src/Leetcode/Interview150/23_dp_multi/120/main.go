package main

import "math"

// 状态为 dfs(i,j)，表示从 (i,j) 出发，移动到 triangle 最后一排，路径上的元素之和的最小值。
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = math.MinInt // 标识没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n-1 {
			return triangle[i][j]
		}
		p := &memo[i][j]
		if *p != math.MinInt {
			return *p
		}
		*p = min(dfs(i+1, j), dfs(i+1, j+1)) + triangle[i][j]
		return *p
	}
	return dfs(0, 0)
}

func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, i+1)
	}
	f[n-1] = triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j, x := range triangle[i] {
			f[i][j] = min(f[i+1][j], f[i+1][j+1]) + x
		}
	}
	return f[0][0]
}

func minimumTotal3(f [][]int) int {
	for i := len(f) - 2; i >= 0; i-- {
		for j := range f[i] {
			f[i][j] += min(f[i+1][j], f[i+1][j+1])
		}
	}
	return f[0][0]
}
