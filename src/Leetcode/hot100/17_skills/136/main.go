package main

func singleNumber(nums []int) int {
	x := nums[0]
	for _, v := range nums[1:] {
		x ^= v
	}
	return x
}
