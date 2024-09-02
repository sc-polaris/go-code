package main

import "slices"

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func removeDuplicates2(nums []int) int {
	return len(slices.Compact(nums))
}

func removeDuplicates3(nums []int) int {
	k := 1
	u := 0
	for _, v := range nums {
		if u < k || nums[u-k] != v {
			nums[u] = v
			u++
		}
	}
	return u
}
