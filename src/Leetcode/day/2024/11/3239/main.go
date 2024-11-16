package main

/*
	给你一个 m x n 的二进制矩阵 grid 。

	如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。

	你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。

	请你返回 最少 翻转次数，使得矩阵 要么 所有行是 回文的 ，要么所有列是 回文的 。
*/

func minFlips(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	diffRow := 0
	for _, row := range grid {
		for j := 0; j < n/2; j++ {
			if row[j] != row[n-1-j] {
				diffRow++
			}
		}
	}

	diffCol := 0
	for j := range grid[0] {
		for i := range grid[:m/2] {
			if grid[i][j] != grid[m-1-i][j] {
				diffCol++
			}
		}
	}

	return min(diffRow, diffCol)
}
