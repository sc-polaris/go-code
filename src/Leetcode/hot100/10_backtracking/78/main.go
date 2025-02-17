package main

import "slices"

// 选或不选
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		// 不选
		dfs(i + 1)

		// 选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1] // 恢复
	}
	dfs(0)
	return
}

// 枚举选哪个
func subsets2(nums []int) (ans [][]int) {
	n := len(nums)
	var path []int
	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, slices.Clone(path))
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
