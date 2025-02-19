package main

import (
	"math"
	"slices"
)

func maxProduct(nums []int) int {
	n := len(nums)
	fMax := make([]int, n)
	fMin := make([]int, n)
	fMax[0], fMin[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		fMax[i] = max(fMax[i-1]*x, fMin[i-1]*x, x)
		fMin[i] = min(fMax[i-1]*x, fMin[i-1]*x, x)
	}
	return slices.Max(fMax)
}

func maxProduct2(nums []int) int {
	ans := math.MinInt
	fMax, fMin := 1, 1
	for _, x := range nums {
		fMax, fMin = max(fMax*x, fMin*x, x), min(fMax*x, fMin*x, x)
		ans = max(ans, fMax)
	}
	return ans
}
