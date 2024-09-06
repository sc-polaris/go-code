package main

// 方法一：哈希表
func containsNearbyDuplicate(nums []int, k int) bool {
	pos := make(map[int]int)
	for i, x := range nums {
		if p, ok := pos[x]; ok && i-p <= k {
			return true
		}
		pos[x] = i
	}
	return false
}

// 方法二：滑动窗口
func containsNearbyDuplicate2(nums []int, k int) bool {
	set := make(map[int]struct{})
	for i, x := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[x]; ok {
			return true
		}
		set[x] = struct{}{}
	}
	return false
}
