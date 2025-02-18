package main

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		mid := (i + j) / 2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
	//return sort.SearchInts(nums, target)
}
