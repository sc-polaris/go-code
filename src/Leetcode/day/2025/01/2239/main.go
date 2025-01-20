package main

/*
	给你一个长度为 n 的整数数组 nums ，请你返回 nums 中最 接近 0 的数字。如果有多个答案，请你返回它们中的 最大值 。
*/

func findClosestNumber(nums []int) int {
	ans := nums[0]
	for _, x := range nums {
		if abs(x) < abs(ans) || abs(x) == abs(ans) && x > 0 {
			ans = x
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
