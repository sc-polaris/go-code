package main

/*
	给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	i, j, di := 0, 0, 0
	for val := 1; val <= n*n; val++ {
		ans[i][j] = val
		x, y := i+dirs[di][0], j+dirs[di][1] // 下一步的位置
		if x < 0 || x >= n || y < 0 || y >= n || ans[x][y] != 0 {
			di = (di + 1) % 4 // 右转 90°
		}
		i += dirs[di][0]
		j += dirs[di][1]
	}
	return ans
}
