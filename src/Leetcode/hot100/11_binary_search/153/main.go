package main

import (
	"sort"
)

func findMin(nums []int) int {
	i := sort.Search(len(nums)-1, func(i int) bool { return nums[i] < nums[len(nums)-1] })
	return nums[i]
}
