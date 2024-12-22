package main

import (
	"cmp"
	"slices"
)

/*
	我们将整数 x 的 权重 定义为按照下述规则将 x 变成 1 所需要的步数：
	· 如果 x 是偶数，那么 x = x / 2
	· 如果 x 是奇数，那么 x = 3 * x + 1
	比方说，x=3 的权重为 7 。因为 3 需要 7 步变成 1 （3 --> 10 --> 5 --> 16 --> 8 --> 4 --> 2 --> 1）。

	给你三个整数 lo， hi 和 k 。你的任务是将区间 [lo, hi] 之间的整数按照它们的权重 升序排序 ，如果大于等于 2 个整数有 相同 的权重，那么按照数字自身的数值 升序排序 。

	请你返回区间 [lo, hi] 之间的整数按权重排序后的第 k 个数。

	注意，题目保证对于任意整数 x （lo <= x <= hi） ，它变成 1 所需要的步数是一个 32 位有符号整数。
*/

/*
	如果 i 是偶数，那么 dfs(i)=dfs(i/2)+1。
	如果 i 是奇数，那么 dfs(i)=dfs(3i+1)+1。进一步地，由于 i 是奇数，3i+1 一定是偶数，所以也可以直接走两步，写成 dfs(i)=dfs((3i+1)/2)+2。
*/

var memo = make(map[int]int)

func dfs(i int) int {
	if i == 1 {
		return 0
	}
	if res, ok := memo[i]; ok {
		return res
	}
	if i&1 == 1 {
		memo[i] = dfs((i*3+1)/2) + 2
	} else {
		memo[i] = dfs(i/2) + 1
	}
	return memo[i]
}

func getKth(lo int, hi int, k int) int {
	nums := make([]int, hi-lo+1)
	for i := range nums {
		nums[i] = i + lo
	}
	slices.SortFunc(nums, func(x, y int) int { return cmp.Or(dfs(x)-dfs(y), x-y) })
	return nums[k-1]
}
