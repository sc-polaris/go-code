package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, sum, l := n+1, 0, 0
	for r, x := range nums {
		sum += x
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if ans <= n {
		return ans
	}
	return 0
}
