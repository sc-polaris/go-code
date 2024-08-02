package main

/*
	给你一个二维 boolean 矩阵 grid 。
	请你返回使用 grid 中的 3 个元素可以构建的 直角三角形 数目，且满足 3 个元素值 都 为 1 。
	注意：
	· 如果 grid 中 3 个元素满足：一个元素与另一个元素在 同一行，同时与第三个元素在 同一列 ，那么这 3 个元素称为一个 直角三角形 。这 3 个元素互相之间不需要相邻。
*/

/*
	套路：有三个顶点，枚举「中间」的直角顶点更容易计算。
	想一想，直角顶点为 (i,j) 的「直角三角形」有多少个？
	设第 i 行有 rowSum 个 1，第 j 列有 colSum 个 1。根据乘法原理，直角顶点为 (i,j) 的「直角三角形」有
						(rowSum−1)⋅(colSum−1)
	个，加到答案中。
*/

func numberOfRightTriangles(grid [][]int) (ans int64) {
	n := len(grid[0])
	colSum := make([]int, n)
	for _, row := range grid {
		for j, x := range row {
			colSum[j] += x
		}
	}

	for _, row := range grid {
		rowSum := 0
		for _, x := range row {
			rowSum += x
		}
		for j, x := range row {
			if x == 1 {
				ans += int64((rowSum - 1) * (colSum[j] - 1))
			}
		}
	}
	return
}
