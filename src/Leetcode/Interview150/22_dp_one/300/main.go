package main

import (
	"slices"
)

func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n)
	var dfs func(int) int
	dfs = func(i int) int {
		p := &memo[i]
		if *p > 0 {
			return *p
		}
		res := 0
		for j, x := range nums[:i] {
			if x < nums[i] {
				res = max(res, dfs(j))
			}
		}
		res++
		*p = res
		return res
	}
	for i := range n {
		ans = max(ans, dfs(i))
	}
	return
}

func lengthOfLIS2(nums []int) (ans int) {
	f := make([]int, len(nums))
	for i, x := range nums {
		for j, y := range nums[:i] {
			if y < x {
				f[i] = max(f[i], f[j])
			}
		}
		f[i]++
	}
	return slices.Max(f)
}
