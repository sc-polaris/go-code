package main

import "slices"

/*
	给你一个整数数组 nums，和一个整数 k 。

	对于每个下标 i（0 <= i < nums.length），将 nums[i] 变成 nums[i] + k 或 nums[i] - k 。

	nums 的 分数 是 nums 中最大元素和最小元素的差值。

	在更改每个下标对应的值之后，返回 nums 的最小 分数 。
*/

func smallestRangeII(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := 1; i < n; i++ {
		mx := max(nums[i-1]+k, nums[n-1]-k)
		mn := min(nums[0]+k, nums[i]-k)
		ans = min(ans, mx-mn)
	}
	return ans
}
