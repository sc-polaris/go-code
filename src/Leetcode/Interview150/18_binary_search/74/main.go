package main

// 二分
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n
	for l < r {
		mid := l + (r-l)/2
		x := matrix[mid/n][mid%n]
		if x == target {
			return true
		}
		if x < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}

// 排除法
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 { // 还有剩余元素
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++ // 这一行剩余元素全部小于 target，排除
		} else {
			j-- // 这一行剩余元素全部大于 target，排除
		}
	}
	return false
}
