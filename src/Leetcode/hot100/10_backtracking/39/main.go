package main

import "slices"

// 选或不选
func combinationSum(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	var path []int
	var dfs func(int, int)
	dfs = func(i, x int) {
		if x == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}
		if i == len(candidates) || x < candidates[i] {
			return
		}

		// 不选
		dfs(i+1, x)
		// 选
		path = append(path, candidates[i])
		dfs(i, x-candidates[i])
		path = path[:len(path)-1]
	}
	dfs(0, target)
	return
}

// 枚举选哪个
func combinationSum2(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	var path []int
	var dfs func(int, int)
	dfs = func(i, x int) {
		if x == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}
		if i == len(candidates) || x < candidates[i] {
			return
		}

		for j := i; j < len(candidates); j++ {
			path = append(path, candidates[j])
			dfs(j, x-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return
}
