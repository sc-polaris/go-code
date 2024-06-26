package main

/*
	给你一个下标从 0 开始的整数数组 nums ，它包含 n 个 互不相同 的正整数。如果 nums 的一个排列满足以下条件，我们称它是一个特别的排列：
	· 对于 0 <= i < n - 1 的下标 i ，要么 nums[i] % nums[i+1] == 0 ，要么 nums[i+1] % nums[i] == 0 。
	请你返回特别排列的总数目，由于答案可能很大，请将它对 10^9 + 7 取余 后返回。
*/

/*
	dfs(S,i) 表示在可以选择的下标集合为 S，上一个选的数下标为 i 时，可以构造出多少个特别排列。
	枚举当前要选的数的下标 j，那么接下来要解决的问题是，在可以选的下标集合为 S\{j}，上一个选的数的下标是 j 时，可以构造出多少个特别队列。
	累加这些方案

*/

func specialPerm(nums []int) (ans int) {
	n := len(nums)
	u := 1<<n - 1
	memo := make([][]int, u)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(s, i int) (res int) {
		if s == 0 { // 找到一个特别排列
			return 1
		}
		p := &memo[s][i]
		if *p != -1 { // 之前计算过
			return *p
		}
		for j, x := range nums {
			if s>>j&1 == 1 && (nums[i]%x == 0 || x%nums[i] == 0) {
				res += dfs(s^(1<<j), j)
			}
		}
		*p = res // 记忆化
		return
	}
	for i := range nums {
		ans += dfs(u^(1<<i), i)
	}
	return ans % 1_000_000_007
}
