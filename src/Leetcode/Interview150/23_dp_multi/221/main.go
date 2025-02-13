package main

// f[i][j] 表示以[i][j]为右下端点的正方形的边长
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	maxSide := 0
	f := make([][]int, m)
	for i, row := range matrix {
		f[i] = make([]int, n)
		for j, x := range row {
			f[i][j] = int(x - '0')
			if f[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if f[i][j] == 1 {
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
				if f[i][j] > maxSide {
					maxSide = f[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}
