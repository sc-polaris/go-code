package main

import "sort"

/*
	设 x=nums[mid] 是我们现在二分取到的数。

	现在需要判断 x 和 target 的位置关系，谁在左边，谁在右边？

	核心思路
	· 如果 x 和 target 在不同的递增段：
		· 如果 target 在第一段（左），x 在第二段（右），说明 x 在 target 右边；
		· 如果 target 在第二段（右），x 在第一段（左），说明 x 在 target 左边。
	· 如果 x 和 target 在相同的递增段：
		· 比较 x 和 target 的大小即可。

	分类讨论
	下面只讨论 x 在 target 右边，或者等于 target 的情况。其余情况 x 一定在 target 左边。
	· 如果 x>nums[n−1]，说明 x 在第一段中，那么 target 也必须在第一段中（否则 x 一定在 target 左边）且 x 必须大于等于 target。
		· 写成代码就是 target > nums[n - 1] && x >= target。
	· 如果 x≤nums[n−1]，说明 x 在第二段中（或者 nums 只有一段），那么 target 可以在第一段，也可以在第二段。
		· 如果 target 在第一段，那么 x 一定在 target 右边。
		· 如果 target 在第二段，那么 x 必须大于等于 target。
		· 写成代码就是 target > nums[n - 1] || x >= target。
	根据这两种情况，去判断 x 和 target 的位置关系，从而不断地缩小 target 所在位置的范围，二分找到 target。
*/

func search(nums []int, target int) int {
	end := nums[len(nums)-1]
	i := sort.Search(len(nums), func(i int) bool {
		x := nums[i]
		if x > end {
			return target > end && x >= target
		}
		return target > end || x >= target
	})
	if nums[i] != target {
		return -1
	}
	return i
}
