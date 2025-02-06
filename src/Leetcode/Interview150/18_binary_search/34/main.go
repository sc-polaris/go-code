package main

import (
	"slices"
	"sort"
)

func searchRange(nums []int, target int) []int {
	search := func(nums []int, target int) int {
		l, r := 0, len(nums)
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}

	i := search(nums, target)
	if i == len(nums) || nums[i] != target {
		return []int{-1, -1}
	}
	j := search(nums, target+1) - 1
	return []int{i, j}

	//i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	//if i == len(nums) || nums[i] != target {
	//	return []int{-1, -1}
	//}
	//j := sort.Search(len(nums), func(i int) bool { return nums[i] >= target+1 }) - 1
	//return []int{i, j}
}

func searchRange2(nums []int, target int) []int {
	start, ok := slices.BinarySearch(nums, target)
	if !ok {
		return []int{-1, -1}
	}
	end := sort.SearchInts(nums, target+1) - 1
	return []int{start, end}
}
