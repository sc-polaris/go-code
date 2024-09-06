package main

// 暴力
func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// hash
func twoSum2(nums []int, target int) []int {
	idx := make(map[int]int)
	for j, x := range nums {
		if i, ok := idx[target-x]; ok {
			return []int{i, j}
		}
		idx[x] = j
	}
	return nil
}
