package main

import "slices"

func rotate(matrix [][]int) {
	// 水平翻转
	//for i := 0; i < n/2; i++ {
	//	matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	//}
	slices.Reverse(matrix)

	// 对角翻转
	for i := range matrix {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
