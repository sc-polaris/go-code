package main

import (
	"math"
	"sort"
)

/*
	f[i][j] 表示从 (i,j) 到 (m-1,n-1) 经过的最少格子数
	1. 从 (i,j) 向右移动一次，到达 (i,k)，问题变成从 (i,k) 到 (m-1,n-1) 经过的最少格子数，即 f[i][j]=f[i][k]+1
	2. 从 (i,j) 向下移动一次，到达 (k,j)，问题变成从 (k,j) 到 (m-1,n-1) 经过的最少格子数，即 f[i][j]=f[k][j]+1
	倒序枚举 i 和 j 计算。答案为 f[0][0]
*/

func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colSts := make([][]pair, n) // 每列的单调栈
	var rowSt []pair
	for i := m - 1; i >= 0; i-- {
		rowSt = rowSt[:0]
		for j := n - 1; j >= 0; j-- {
			colSt := colSts[j]
			if i < m-1 || j < n-1 {
				mn = math.MaxInt
			}
			if g := grid[i][j]; g > 0 { // 可以向右/向下跳
				k := sort.Search(len(rowSt), func(k int) bool { return rowSt[k].i <= g+j })
				if k < len(rowSt) {
					mn = rowSt[k].x
				}
				k = sort.Search(len(colSt), func(k int) bool { return colSt[k].i <= g+i })
				if k < len(colSt) {
					mn = min(mn, colSt[k].x)
				}
			}
			if mn < math.MaxInt {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(rowSt) > 0 && mn <= rowSt[len(rowSt)-1].x {
					rowSt = rowSt[:len(rowSt)-1]
				}
				rowSt = append(rowSt, pair{mn, j})
				for len(colSt) > 0 && mn <= colSt[len(colSt)-1].x {
					colSt = colSt[:len(colSt)-1]
				}
				colSts[j] = append(colSt, pair{mn, i})
			}
		}
	}

	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt {
		return -1
	}
	return
}

func main() {

}
