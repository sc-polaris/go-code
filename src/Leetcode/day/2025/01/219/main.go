package main

/*
	给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且
	abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	s := make(map[int]struct{})
	for i, x := range nums {
		if i > k {
			delete(s, nums[i-k-1])
		}
		if _, ok := s[x]; ok {
			return true
		}
		s[x] = struct{}{}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	last := make(map[int]int)
	for i, x := range nums {
		if j, ok := last[x]; ok && i-j <= k {
			return true
		}
		last[x] = i
	}
	return false
}
