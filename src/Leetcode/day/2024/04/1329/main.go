package main

import "slices"

/*
	示例 1:
	Before：			After：
	3 3 1 1				1 1 1 1
	2 2 1 2				1 2 2 2
	1 1 1 2				1 2 3 3

	一条对角线上的坐标（行列下标）满足什么性质？
	比如示例 1 的对角线 3, 1, 2，坐标分别为 (0,1), (1,2), (2,3)，它们都满足：行下标减列下标等于一个定值 k = -1。

	考虑从右上第一条对角线（上图中的 (0,3)）开始，一条一条地排序，直到左下最后一条对角线（上图中的 (2,0)）结束。

	设坐标为 (i,j)，设 k = i - j。
	· 第一条对角线上只有一个点，坐标为 (0,n-1)，其 k = 1 - n。
	· 最后一条对角线上也只有一个点，坐标为 (m-1,0)，其 k = m - 1。
	· 所以枚举对角线，就是枚举 k 从 1 - n 到 m - 1。

	对于同一条对角线，直到行下标 i 就知道列下标 j = i - k。
	· i 的最小值：令等式 k = i - j 中的 j = 0，可得 i = k，但 i 必须是非负数，所以 i 最小为 max(k,0)。
	· i 的最大值：令等式 k = i - j 中的 j = n - 1，可得 i = k + n - 1，但 i 至多为 m - 1，所以 i 最大为 min(k+n-1,m-1)。
	· 枚举 i，范围为左闭右开去见 [max(k,0),min(k+n,m)]。
	· 把 mat[i][i-k] 加入一个数组，把数组从小到大排序后，再依次填入 mat[i][i-k]，即完成了这条对角线的排序。
*/

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	arr := make([]int, min(m, n))
	for k := 1 - n; k < m; k++ { // k = i - j
		a := arr[:0]
		minI := max(k, 0)
		maxI := min(k+n, m)
		for i := minI; i < maxI; i++ {
			a = append(a, mat[i][i-k])
		}
		slices.Sort(a)
		for i := minI; i < maxI; i++ {
			mat[i][i-k] = a[i-minI]
		}
	}
	return mat
}
