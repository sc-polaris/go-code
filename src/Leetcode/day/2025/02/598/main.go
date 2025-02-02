package main

/*
	给你一个 m x n 的矩阵 M 和一个操作数组 op 。矩阵初始化时所有的单元格都为 0 。ops[i] = [ai, bi] 意味着当所有的 0 <= x < ai 和
	0 <= y < bi 时， M[x][y] 应该加 1。

	在 执行完所有操作后 ，计算并返回 矩阵中最大整数的个数 。
*/

func maxCount(m int, n int, ops [][]int) int {
	minA, minB := m, n
	for _, op := range ops {
		minA = min(minA, op[0])
		minB = min(minB, op[1])
	}
	return minA * minB
}
