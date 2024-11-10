package main

import "sort"

/*
	给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。

	请你找出并返回只出现一次的那个数。

	你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
*/

/*
	题目有两个已知条件：
	1. 数组是有序的。
	2. 除了一个数出现一次外，其余每个数都出现两次。
	第二个条件意味着，数组的长度一定是奇数。
	第一个条件意味着，出现两次的数，必然相邻，不可能出现 1,2,1 这样的顺序。
	这也意味着，只出现一次的那个数，一定位于偶数下标上。

	这启发我们去检查偶数下标 2k。
	示例 1 的 nums=[1,1,2,3,3,4,4,8,8]：
	· 如果 nums[2k]==nums[2k+1]，说明只出现一次的数的下标 > 2k。
	· 如果 nums[2k]!=nums[2k+1]，说明只出现一次的数的下标 ≤ 2k。
	也就是说，随着 k 的变大，不等式 nums[2k]!=nums[2k+1] 越可能满足，有单调性，可以二分。
*/

func singleNonDuplicate(nums []int) int {
	i, j := 0, len(nums)/2
	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[h*2] == nums[h*2+1] {
			i = h + 1
		} else {
			j = h
		}
	}

	return nums[2*i]
}

func singleNonDuplicate2(nums []int) int {
	k := sort.Search(len(nums)/2, func(k int) bool {
		return nums[k*2] != nums[k*2+1]
	})
	return nums[2*k]
}
