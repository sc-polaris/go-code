package main

import "math"

/*
	给你一个下标从 0 开始、长度为 n 的整数数组 nums ，以及整数 indexDifference 和整数 valueDifference 。

	你的任务是从范围 [0, n - 1] 内找出  2 个满足下述所有条件的下标 i 和 j ：
	· abs(i - j) >= indexDifference 且
	· abs(nums[i] - nums[j]) >= valueDifference
	返回整数数组 answer。如果存在满足题目要求的两个下标，则 answer = [i, j] ；否则，answer = [-1, -1] 。如果存在多组可供选择的下标对，只需要返回其中任意一组即可。

	注意：i 和 j 可能 相等 。
*/

// O(n^2(

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := range nums {
		for j := i; j < len(nums); j++ {
			if j-i >= indexDifference && int(math.Abs(float64(nums[i]-nums[j]))) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func findIndices2(nums []int, indexDifference int, valueDifference int) []int {
	maxIdx, minIdx := 0, 0
	for j := indexDifference; j < len(nums); j++ {
		i := j - indexDifference
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		} else if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[maxIdx]-nums[j] >= valueDifference {
			return []int{maxIdx, j}
		}
		if nums[j]-nums[minIdx] >= valueDifference {
			return []int{minIdx, j}
		}
	}
	return []int{-1, -1}
}
