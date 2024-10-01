package main

import (
	"math"
	"slices"
)

/*
	给你一个长度为 n 的整数数组 nums 和一个 正 整数 k 。

	一个子序列的 能量 定义为子序列中 任意 两个元素的差值绝对值的 最小值 。

	请你返回 nums 中长度 等于 k 的 所有 子序列的 能量和 。

	由于答案可能会很大，将答案对 10^9 + 7 取余 后返回。
*/

/*
	设计一个函数 dfs(i,j,k,mi)，表示当前处理到第 i 个元素，上一个选取的是第 j 个元素，还需要选取 k 个元素，
	当前的最小差值为 mi 时，能量和的值。那么答案就是 dfs(0,n,k,+∞)。（若上一个选取的是第 n 个元素，表示之前
	没有选取过元素）。
*/

func sumOfPowers(nums []int, k int) int {
	const mod = 1e9 + 7
	slices.Sort(nums)
	n := len(nums)
	f := make(map[int]int)
	var dfs func(i, j, k, mi int) int
	dfs = func(i, j, k, mi int) int {
		if i >= n {
			if k == 0 {
				return mi
			}
			return 0
		}
		if n-i < k {
			return 0
		}
		// 记忆化搜索保存
		key := mi<<18 | (i << 12) | (j << 6) | k
		if v, ok := f[key]; ok {
			return v
		}
		ans := dfs(i+1, j, k, mi) // 不选第 i 个元素
		if j == n {
			ans += dfs(i+1, i, k-1, mi)
		} else {
			ans += dfs(i+1, i, k-1, min(mi, nums[i]-nums[j]))
		}
		ans %= mod
		f[key] = ans
		return ans
	}
	return dfs(0, n, k, math.MaxInt)
}
