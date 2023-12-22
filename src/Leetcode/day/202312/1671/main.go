package main

import (
	"sort"
)

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	suf := make([]int, n)
	var g []int
	for i := n - 1; i > 0; i-- {
		x := nums[i]
		j := sort.SearchInts(g, x)
		if j == len(g) {
			g = append(g, x)
		} else {
			g[j] = x
		}
		suf[i] = j + 1 // 从 nums[i] 开始的最长严格递减子序列的长度
	}

	mx := 0
	g = g[:0]
	for i, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) {
			g = append(g, x)
		} else {
			g[j] = x
		}
		pre := j + 1 // 在 nums[j] 结束的最长严格递增子序列的长度
		if pre >= 2 && suf[i] >= 2 {
			mx = max(mx, pre+suf[i]-1) // 减去重复的 nums[i]
		}
	}
	return n - mx
}
