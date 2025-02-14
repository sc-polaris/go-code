package main

import "slices"

func rotate(matrix [][]int) {
	slices.Reverse(matrix)

	for i := range matrix {
		for j := range i {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
