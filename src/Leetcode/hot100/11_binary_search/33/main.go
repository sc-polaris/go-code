package main

import "sort"

/*
	如果 x 和 target 在不同的递增段：
		如果 target 在第一段（左），x 在第二段（右），说明 x 在 target 右边；
		如果 target 在第二段（右），x 在第一段（左），说明 x 在 target 左边。
	如果 x 和 target 在相同的递增段：
		比较 x 和 target 的大小即可。
*/

func search(nums []int, target int) int {
	end := nums[len(nums)-1]
	// 只讨论 x 在 target 右边，或者等于 target 的情况。其余情况 x 一定在 target 左边。
	i := sort.Search(len(nums), func(i int) bool {
		x := nums[i]
		if x > end { // 第一段
			return target > end && x >= target
		}
		return target > end || x >= target
	})
	if nums[i] != target {
		return -1
	}
	return i
}
