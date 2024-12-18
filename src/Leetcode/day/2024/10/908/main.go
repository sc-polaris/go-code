package main

import "slices"

/*
	给你一个整数数组 nums，和一个整数 k 。

	在一个操作中，您可以选择 0 <= i < nums.length 的任何索引 i 。将 nums[i] 改为 nums[i] + x ，其中 x 是一个范围为 [-k, k] 的任意整数。
	对于每个索引 i ，最多 只能 应用 一次 此操作。

	nums 的 分数 是 nums 中最大和最小元素的差值。

	在对  nums 中的每个索引最多应用一次上述操作后，返回 nums 的最低 分数 。
*/

func smallestRangeI(nums []int, k int) int {
	return max(slices.Max(nums)-slices.Min(nums)-k*2, 0)
}
