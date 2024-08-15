package main

import "math"

/*
	给你一个由 正整数 组成、大小为 m x n 的矩阵 grid。你可以从矩阵中的任一单元格移动到另一个位于正下方或正右侧的任意单元格（不必相邻）。
	从值为 c1 的单元格移动到值为 c2 的单元格的得分为 c2 - c1 。

	你可以从 任一 单元格开始，并且必须至少移动一次。

	返回你能得到的 最大 总得分。
*/

/*
	脑筋急转弯的二维前缀和

	定义 f[i+1][j+1] 表示左上角在 (0,0)，右下角在 (i,j) 的子矩阵的最小值。疑似二维前缀和，f[i+1][j+1] 可以递推计算：
				f[i+1][j+1] = min(f[i+1][j],f[i][j+1],grid[i][j])
	注意题目要求至少移动一次，也就是起点与终点不能重合。如果终点在 (i,j)，那么起点的海拔高度最小值为
				min(f[i+1][j],f[i][j+1])
	终点与起点的海拔高度之差为
				grid[i][j]−min(f[i+1][j],f[i][j+1])
*/

func maxScore(grid [][]int) int {
	ans := math.MinInt
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for j := range f[0] {
		f[0][j] = math.MaxInt
	}
	for i, row := range grid {
		f[i+1] = make([]int, n+1)
		f[i+1][0] = math.MaxInt
		for j, x := range row {
			mn := min(f[i+1][j], f[i][j+1])
			ans = max(ans, x-mn)
			f[i+1][j+1] = min(mn, x)
		}
	}
	return ans
}

// 优化 也可以维护每列的最小值 colMin，这样空间复杂度更小。

func maxScore2(grid [][]int) int {
	ans := math.MinInt
	colMin := make([]int, len(grid[0]))
	for i := range colMin {
		colMin[i] = math.MaxInt
	}
	for _, row := range grid {
		preMin := math.MaxInt // colMin[0..j] 的最小值
		for j, x := range row {
			ans = max(ans, x-min(preMin, colMin[j]))
			colMin[j] = min(colMin[j], x)
			preMin = min(preMin, colMin[j])
		}
	}
	return ans
}
