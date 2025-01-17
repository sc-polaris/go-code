package main

import "math"

/*
	给你一个 非负 整数数组 nums 和一个整数 k 。

	如果一个数组中所有元素的按位或运算 OR 的值 至少 为 k ，那么我们称这个数组是 特别的 。

	请你返回 nums 中 最短特别非空子数组
	的长度，如果特别子数组不存在，那么返回 -1 。
*/

// minimumSubarrayLength LogTrick
func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		if x >= k {
			return 1
		}
		for j := i - 1; j >= 0 && nums[j]|x != nums[j]; j-- {
			nums[j] |= x
			if nums[j] >= k {
				ans = min(ans, i-j+1)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// minimumSubarrayLength2 滑动窗口+栈
func minimumSubarrayLength2(nums []int, k int) int {
	ans := math.MaxInt
	var left, bottom, rightOr int
	for right, x := range nums {
		rightOr |= x
		for left <= right && nums[left]|rightOr >= k {
			ans = min(ans, right-left+1)
			left++
			if bottom < left {
				for i := right - 1; i >= left; i-- {
					nums[i] |= nums[i+1]
				}
				bottom = right
				rightOr = 0
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
