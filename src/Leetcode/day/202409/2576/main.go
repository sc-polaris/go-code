package main

import (
	"slices"
	"sort"
)

/*
	给你一个下标从 0 开始的整数数组 nums 。

	一开始，所有下标都没有被标记。你可以执行以下操作任意次：
	· 选择两个 互不相同且未标记 的下标 i 和 j ，满足 2 * nums[i] <= nums[j] ，标记下标 i 和 j 。
	请你执行上述操作任意次，返回 nums 中最多可以标记的下标数目。
*/

/*
	方法一：二分答案
	如果 2⋅nums≤nums[j]，则称 nums[i] 与 nums[j] 匹配。
	如果可以匹配 k 对，那么也可以匹配小于 k 对，去掉一些数对即可做到。
	如果无法匹配 k 对，那么也无法匹配大于 k 对（反证法）。
	所以 k 越大，越无法选出 k 个能匹配的数对。有单调性，就可以二分答案。

	现在问题变成：能否从 nums 中选出 k 个能匹配的数对？
	结论：从小到大排序后，如果存在 k 对匹配，那么一定可以让最小的 k 个数与最大的 k 个数匹配。

	证明：假设不是最小的 k 个数与最大的 k 个数匹配，那么我们总是可以把 nums[i] 替换成比它小的且不在匹配中的数，这仍然是
	匹配的；同理，把 nums[j] 替换成比它大的且不在匹配中的数，这仍然是匹配的。所以如果存在 k 对匹配，那么一定可以让最小的
	k 个数和最大的 k 个数匹配。

*/

func maxNumOfMarkedIndices(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	pairs := sort.Search(n/2, func(k int) bool {
		k++
		for i, x := range nums[:k] {
			if x*2 > nums[n-k+i] {
				return true
			}
		}
		return false
	})
	return pairs * 2 // 最多匹配 pairs 对，有 pairs * 2 个数
}

// 方法二：同向双指针

func maxNumOfMarkedIndices2(nums []int) int {
	slices.Sort(nums)
	i := 0
	for _, x := range nums[(len(nums)+1)/2:] {
		if nums[i]*2 <= x { // 找到一个匹配
			i++
		}
	}
	return i * 2
}
