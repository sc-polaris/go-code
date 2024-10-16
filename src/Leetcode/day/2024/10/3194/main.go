package main

import (
	"math"
	"slices"
)

/*
	你有一个初始为空的浮点数数组 averages。另给你一个包含 n 个整数的数组 nums，其中 n 为偶数。

	你需要重复以下步骤 n / 2 次：
	· 从 nums 中移除 最小 的元素 minElement 和 最大 的元素 maxElement。
	· 将 (minElement + maxElement) / 2 加入到 averages 中。
	返回 averages 中的 最小 元素。
*/

func minimumAverage(nums []int) float64 {
	slices.Sort(nums)
	ans := math.MaxInt
	for i, n := 0, len(nums); i < n/2; i++ {
		ans = min(ans, nums[i]+nums[n-1-i])
	}
	return float64(ans) / 2
}
