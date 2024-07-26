package main

import (
	"math"
	"slices"
)

/*
	给你一个 正 整数数组 nums 。

	将 nums 分成两个数组：nums1 和 nums2 ，并满足下述条件：
	· 数组 nums 中的每个元素都属于数组 nums1 或数组 nums2 。
	· 两个数组都 非空 。
	· 分区值 最小 。
	分区值的计算方法是 |max(nums1) - min(nums2)| 。

	其中，max(nums1) 表示数组 nums1 中的最大元素，min(nums2) 表示数组 nums2 中的最小元素。

	返回表示分区值的整数。
*/

func findValueOfPartition(nums []int) int {
	slices.Sort(nums)
	ans := math.MaxInt
	for i := 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
	}
	return ans
}
