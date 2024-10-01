package main

import "math"

/*
	给你一个下标从 0 开始的整数数组 nums 和一个正整数 x 。

	你 一开始 在数组的位置 0 处，你可以按照下述规则访问数组中的其他位置：

	如果你当前在位置 i ，那么你可以移动到满足 i < j 的 任意 位置 j 。
	对于你访问的位置 i ，你可以获得分数 nums[i] 。
	如果你从位置 i 移动到位置 j 且 nums[i] 和 nums[j] 的 奇偶性 不同，那么你将失去分数 x 。
	请你返回你能得到的 最大 得分之和。

	注意 ，你一开始的分数为 nums[0] 。
*/

func maxScore(nums []int, x int) int64 {
	res := int64(nums[0])
	dp := [2]int64{math.MinInt32, math.MinInt32}
	dp[nums[0]%2] = int64(nums[0])
	for i := 1; i < len(nums); i++ {
		parity := nums[i] % 2
		cur := max(dp[parity]+int64(nums[i]), dp[1-parity]-int64(x)+int64(nums[i]))
		res = max(res, cur)
		dp[parity] = max(dp[parity], cur)
	}
	return res
}
