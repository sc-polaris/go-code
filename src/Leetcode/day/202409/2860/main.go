package main

import "slices"

/*
	给你一个下标从 0 开始、长度为 n 的整数数组 nums ，其中 n 是班级中学生的总数。班主任希望能够在让所有学生保持开心的情况下
	选出一组学生：

	如果能够满足下述两个条件之一，则认为第 i 位学生将会保持开心：
	1. 这位学生被选中，并且被选中的学生人数 严格大于 nums[i] 。
	2. 这位学生没有被选中，并且被选中的学生人数 严格小于 nums[i] 。

	返回能够满足让所有学生保持开心的分组方法的数目。
*/

/*
	假设恰好选 k 个学生，那么：
	1. 所有 nums[i]<k 的学生都要选；
	2. 所有 nums[i]>k 的学生都不能选；
	3. 不能出现 nums[i]=k 的情况，因为每个学生只有选或不选两种可能。
*/

func countWays(nums []int) (ans int) {
	slices.Sort(nums)
	if nums[0] > 0 { // 一个学生都不选
		ans = 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < i && i < nums[i] {
			ans++
		}
	}
	return ans + 1 // 一定可以都选
}
