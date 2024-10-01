package main

import "strconv"

func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		for _, row := range grid {
			ans[j] = max(ans[j], len(strconv.Itoa(row[j])))
		}
	}
	return ans
}

/*
	由于数字的绝对值越大，数字的长度就越长，所以只需要对每一列的最小值或最大值求长度。
	由于负数中的负号也算一个长度，我们可以取 max(mx,-10*mn) 的长度作为答案
	避免乘法溢出，取 max(mx/10,-mn) 的长度加 1 作为答案。此时要把 0 的长度视为 0。
*/

func findColumnWidth2(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		mn, mx := 0, 0
		for _, row := range grid {
			mn = min(mn, row[j])
			mx = max(mx, row[j])
		}
		xLen := 1
		for x := max(mx/10, -mn); x > 0; x /= 10 {
			xLen++
		}
		ans[j] = xLen
	}
	return ans
}
