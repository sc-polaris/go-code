package main

func swapColors(nums []int, target int) (idx int) {
	for i, c := range nums {
		if c == target {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}
	return idx
}

func sortColors(nums []int) {
	c0 := swapColors(nums, 0)
	swapColors(nums[c0:], 1)
}
