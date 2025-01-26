package main

import "slices"

/*
	给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

	candidates 中的每个数字在每个组合中只能使用 一次 。

	注意：解集不能包含重复的组合。
*/

// 选或不选 combinationSum2
func combinationSum2(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	n := len(candidates)
	var path []int
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 没有可选的数字
		if i == n {
			return
		}

		x := candidates[i]
		if left < x {
			return
		}

		// 选 x
		path = append(path, x)
		dfs(i+1, left-x)
		path = path[:len(path)-1] // 恢复现场

		// 不选 x，那么后面所有等于 x 的数都不选
		// 如果不跳过这些，会导致「选 x 不选 x'」和「不选 x 选 x'」这两种情况都会加到 ans 中，这就重复了
		i++
		for i < n && candidates[i] == x {
			i++
		}
		dfs(i, left)
	}
	dfs(0, target)
	return
}

// 枚举选哪个 combinationSum22
func combinationSum22(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	var path []int
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}

		// 在 [i, len(candidates)-1] 中选一个 candidates【i]
		for j := i; j < len(candidates) && candidates[j] <= left; j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			path = append(path, candidates[j])
			dfs(j+1, left-candidates[j])
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0, target)
	return
}
