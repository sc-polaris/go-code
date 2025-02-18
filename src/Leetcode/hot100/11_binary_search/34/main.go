package main

import "sort"

func searchRange(nums []int, target int) []int {
	i := sort.SearchInts(nums, target)
	if i == len(nums) || nums[i] != target {
		return []int{-1, -1}
	}
	j := sort.SearchInts(nums, target+1) - 1
	return []int{i, j}
}
