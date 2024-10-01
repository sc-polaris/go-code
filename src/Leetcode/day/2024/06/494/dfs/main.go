package dfs

import "math"

/*
	给你一个非负整数数组 nums 和一个整数 target 。
	向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
	· 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
	返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
*/

/*
	设 nums 的元素和为 s，添加正号的元素之和为 p，添加负号的元素（绝对值）之和为 q，那么有
								p + q = s
								p - q = target
	解得
								p = (s+target) / 2
								q = (s-target) / 2
	· 如果 target ≥ 0，那么取 q = (s-target)/2 可以得到更小的背包容量。
	· 如果 target < 0，那么取 p = (s+target)/2 可以得到更小的背包容量。
	综上所述，取
								(s - |target|) / 2
	作为 0-1 背包的背包容量是最优的。
*/

func findTargetSumWays(nums []int, target int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= int(math.Abs(float64(target)))
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2

	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		p := &memo[i][c]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if c < nums[i] {
			return dfs(i-1, c) // 只能不选
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i]) // 不选 + 选
	}
	return dfs(n-1, m)
}
