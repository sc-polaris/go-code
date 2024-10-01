package main

import "slices"

/*
	给你一个下标从 1 开始、大小为 m x n 的整数矩阵 mat，你可以选择任一单元格作为 起始单元格 。

	从起始单元格出发，你可以移动到 同一行或同一列 中的任何其他单元格，但前提是目标单元格的值 严格大于 当前单元格的值。

	你可以多次重复这一过程，从一个单元格移动到另一个单元格，直到无法再进行任何移动。

	请你找出从某个单元开始访问矩阵所能访问的 单元格的最大数量 。

	返回一个表示可访问单元格最大数量的整数。
*/

func maxIncreasingCells(mat [][]int) int {
	type pair struct{ x, y int }
	g := make(map[int][]pair)
	for i, row := range mat {
		for j, x := range row {
			g[x] = append(g[x], pair{i, j}) // 相同元素放在一组，统计位置
		}
	}

	keys := make([]int, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	rowMax := make([]int, len(mat))
	colMax := make([]int, len(mat[0]))
	for _, x := range keys {
		pos := g[x]
		// 先把所有 f 的值都算出来，再更新 rowMax 和 colMax
		fs := make([]int, len(pos))
		for i, p := range pos {
			fs[i] = max(rowMax[p.x], colMax[p.y]) + 1
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], fs[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], fs[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return slices.Max(rowMax)
}
