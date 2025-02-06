package main

import "slices"

/*
	给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
*/

func permuteUnique(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	var path []int
	onPath := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j, on := range onPath {
			if on || j > 0 && nums[j] == nums[j-1] && !onPath[j-1] {
				continue
			}
			path = append(path, nums[j])
			onPath[j] = true
			dfs(i + 1)
			onPath[j] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
