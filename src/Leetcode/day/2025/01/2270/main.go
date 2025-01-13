package main

/*
	给你一个下标从 0 开始长度为 n 的整数数组 nums 。
	如果以下描述为真，那么 nums 在下标 i 处有一个 合法的分割 ：
	· 前 i + 1 个元素的和 大于等于 剩下的 n - i - 1 个元素的和。
	· 下标 i 的右边 至少有一个 元素，也就是说下标 i 满足 0 <= i < n - 1 。
	请你返回 nums 中的 合法分割 方案数。
*/

func waysToSplitArray(nums []int) (ans int) {
	total := 0
	for _, x := range nums {
		total += x
	}

	s := 0
	for _, x := range nums[:len(nums)-1] {
		s += x
		if s*2 >= total {
			ans++
		}
	}
	return
}
