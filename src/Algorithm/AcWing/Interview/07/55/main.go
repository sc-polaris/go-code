package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1], 0) + nums[i]
		res = max(res, nums[i])
	}

	return res
}
