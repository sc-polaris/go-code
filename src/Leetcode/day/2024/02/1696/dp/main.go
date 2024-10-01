package main

import "slices"

func maxResult(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = slices.Max(f[max(i-k, 0):i]) + nums[i]
	}
	return f[n-1]
}

func main() {

}
