package main

import "slices"

// 方法一：选或不选 combinationSum
func combinationSum(candidates []int, target int) (ans [][]int) {
	var path []int
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}

		if i == len(candidates) || left < 0 {
			return
		}
		// 不选
		dfs(i+1, left)
		// 选
		path = append(path, candidates[i])
		dfs(i, left-candidates[i])
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0, target)
	return
}

// 方法一：选或不选 剪枝优化 combinationSum
func combinationSum2(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	var path []int
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}

		if i == len(candidates) || left < candidates[i] {
			return
		}
		// 不选
		dfs(i+1, left)
		// 选
		path = append(path, candidates[i])
		dfs(i, left-candidates[i])
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0, target)
	return
}

// 方法二：枚举选哪个
func combinationSum3(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	var path []int
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}
		if left < candidates[i] {
			return
		}
		for j := i; j < len(candidates); j++ {
			path = append(path, candidates[j])
			dfs(j, left-candidates[j])
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0, target)
	return
}

/*
	方法三：完全背包预处理 + 可行性剪枝
	例如 candidates=[2,4,6,8,10] 都是偶数，但 target=11 是奇数，这种情况我们在一开始递归时，就应当判断出无解，
	不再继续向下递归。
	怎么判断？我们可以用完全背包预处理出下标在 [0,i] 中的 candidates 元素之和能否为 j，记作 f[i+1][j]。
	如果递归中的 left 不在可以组合得到的数字中，则可以直接返回。
	这一做法可以保证我们是在往正确的方向一步步递归前进的。只要题目保证方案数不超过 150，即使 target=1000 也能搞定。
*/

func combinationSum4(candidates []int, target int) (ans [][]int) {
	n := len(candidates)
	// 完全背包
	f := make([][]bool, n+1)
	f[0] = make([]bool, target+1)
	f[0][0] = true
	for i, x := range candidates {
		f[i+1] = make([]bool, target+1)
		for j, b := range f[i] {
			f[i+1][j] = b || j >= x && f[i+1][j-x]
		}
	}

	var path []int
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			// 找到一个合法组合
			ans = append(ans, slices.Clone(path))
			return
		}

		// 无法用下标在 [0, i] 中的数字组合出 left
		if left < 0 || !f[i+1][left] {
			return
		}

		// 不选
		dfs(i-1, left)

		// 选
		path = append(path, candidates[i])
		dfs(i, left-candidates[i])
		path = path[:len(path)-1]
	}

	// 倒着递归，这样参数符合 f 数组的定义
	dfs(n-1, target)
	return
}
