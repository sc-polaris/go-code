package main

import "sort"

func maximumCount(nums []int) int {
	a, b := 0, 0
	for _, x := range nums {
		if x < 0 {
			a++
		} else if x > 0 {
			b++
		}
	}
	return max(a, b)
}

func maximumCount2(nums []int) int {
	// 找到第一个 >= 0 的下标
	x := sort.SearchInts(nums, 0)
	// 第一个 > 0 的位置，等价于第一个 >= 1 的位置
	pos := len(nums) - sort.SearchInts(nums, 1)
	return max(x, pos)
}
