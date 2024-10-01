package main

import "math"

/*
	三元组的题目，通常枚举中间的数
	枚举 nums[j], 我们需要求出 j 左边元素的最小值和右边元素的最小值
	可以用递推 定义 suf[i] 表示从 nums[i] 到 nums[n-1] 的最小值（后缀最小值）, 则有
							suf[i] = min(suf[i+1],nums[i])
	前缀最小值同理
*/

func minimumSum(nums []int) int {
	n := len(nums)
	suf := make([]int, n) // 后缀最小值
	suf[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans := math.MaxInt
	pre := nums[0] // 前缀最小值
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] { // 山形
			ans = min(ans, pre+nums[j]+suf[j+1]) // 更新答案
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}

	return ans
}
