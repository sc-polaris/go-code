package main

import "slices"

/*
	给你两个整数数组 nums1 和 nums2。

	从 nums1 中移除两个元素，并且所有其他元素都与变量 x 所表示的整数相加。如果 x 为负数，则表现为元素值的减少。

	执行上述操作后，nums1 和 nums2 相等 。当两个数组中包含相同的整数，并且这些整数出现的频次相同时，两个数组 相等 。

	返回能够实现数组相等的 最小 整数 x 。
*/

// 由于只能移除两个元素，所以 nums1 的前三小元素必定有一个是保留下来的，我们可以枚举保留下来的最小元素是 nums1[0] 还是 nums1[1] 还是 nums1 2]。

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	// 枚举保留 nums1[2] 或者 nums1[1] 或者 nums1[0]
	// 倒着枚举是因为 nums1[i] 越大答案越小，第一个满足的就是答案
	for i := 2; i > 0; i-- {
		x := nums2[0] - nums1[i]
		// 在 {nums1[i] + x} 中找子序列 nums2
		j := 0
		for _, v := range nums1[i:] {
			if nums2[j] == v+x {
				j++
				// nums2 是 {nums1[i] + x} 的子序列
				if j == len(nums2) {
					return x
				}
			}
		}
	}
	// 题目保证答案一定存在
	return nums2[0] - nums1[0]
}
