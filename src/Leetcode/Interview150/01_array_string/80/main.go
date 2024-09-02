package main

func removeDuplicates(nums []int) int {
	k := 2
	u := 0
	for _, v := range nums {
		if u < k || nums[u-k] != v {
			nums[u] = v
			u++
		}
	}
	return u
}
