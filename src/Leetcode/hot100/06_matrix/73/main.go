package main

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}

// 使用两个标记变量 用矩阵的第一行和第一列代替方法一中的两个标记数组
func setZeroes2(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	r0, c0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			r0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			c0 = true
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if r0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if c0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
