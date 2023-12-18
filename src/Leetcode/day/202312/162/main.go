package main

import (
	"fmt"
	"sort"
)

func findPeakElement(nums []int) int {
	return sort.Search(len(nums)-1, func(i int) bool { return nums[i] > nums[i+1] })
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
