package main

import "slices"

/*
	给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。

	矩阵中坐标 (a, b) 的 值 可由对所有满足 0 <= i <= a < m 且 0 <= j <= b < n 的
	元素 matrix[i][j]（下标从 0 开始计数）执行异或运算得到。

	请你找出 matrix 的所有坐标中第 k 大的值（k 的值从 1 开始计数）。
*/

/*
	一个数异或自己等于 0，s[i+1][j] ^ s[i][j+1] 会导致 s[i][j] 被抵消掉
*/

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	a := make([]int, 0, m*n) // 预分配空间
	s := make([][]int, m+1)
	for i := range s {
		s[i] = make([]int, n+1)
	}
	for i, row := range matrix {
		for j, x := range row {
			s[i+1][j+1] = s[i+1][j] ^ s[i][j+1] ^ s[i][j] ^ x
		}
		a = append(a, s[i+1][1:]...)
	}
	slices.Sort(a)
	return a[len(a)-k]
}
