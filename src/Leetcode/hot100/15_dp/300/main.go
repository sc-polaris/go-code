package main

import (
	"slices"
	"sort"
)

func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		p := &memo[i]
		if *p > 0 {
			return *p
		}
		defer func() { *p = res }()
		for j, x := range nums[:i] {
			if x < nums[i] {
				res = max(res, dfs(j))
			}
		}
		res++
		return
	}
	for i := range nums {
		ans = max(ans, dfs(i))
	}
	return
}

func lengthOfLIS2(nums []int) (ans int) {
	f := make([]int, len(nums))
	for i, x := range nums {
		for j, y := range nums {
			if y < x {
				f[i] = max(f[i], f[j])
			}
		}
		f[i]++
	}
	return slices.Max(f)
}

// 贪心 + 二分查找
func lengthOfLIS3(nums []int) (ans int) {
	var g []int
	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) { // >= x 的 g[j]不存在
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}

func lengthOfLIS4(nums []int) (ans int) {
	g := nums[:0] // 原地修改
	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) { // >= x 的 g[j]不存在
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}
