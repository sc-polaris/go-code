package main

import "slices"

/*
	给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

	解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
*/

// 选或不选
func subsetsWithDup(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 选 x
		x := nums[i]
		path = append(path, x)
		dfs(i + 1)
		path = path[:len(path)-1] // 恢复现场

		// 不选 x 跳过所有等于 x 的
		i++
		for i < n && nums[i] == x {
			i++
		}
		dfs(i)
	}
	dfs(0)
	return
}

// 枚举
func subsetsWithDup2(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, slices.Clone(path))

		for j := i; j < n; j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0)
	return
}
