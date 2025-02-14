package main

import (
	"math"
	"slices"
)

func maxSubArray(nums []int) int {
	ans := math.MinInt
	minPreS, preS := 0, 0
	for _, x := range nums {
		preS += x
		ans = max(ans, preS-minPreS)
		minPreS = min(minPreS, preS)
	}
	return ans
}

// f[i] 表示以 nums[i] 结尾的最大子数组和。
func maxSubArray2(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1], 0) + nums[i]
	}
	return slices.Max(f)
}

func maxSubArray3(nums []int) int {
	ans := math.MinInt
	f := 0
	for _, x := range nums {
		f = max(f, 0) + x
		ans = max(ans, f)
	}
	return ans
}
