package main

/*
	给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，它们分别含有 n 和 m 个元素。请你计算以下两个数值：
	· answer1：使得 nums1[i] 在 nums2 中出现的下标 i 的数量。
	· answer2：使得 nums2[i] 在 nums1 中出现的下标 i 的数量。
	返回 [answer1, answer2]。
*/

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]int)
	for _, v := range nums1 {
		set1[v] = 1
	}
	set2 := make(map[int]int)
	for _, v := range nums2 {
		set2[v] = 1
	}

	ans := [2]int{}
	for _, v := range nums1 {
		ans[0] += set2[v]
	}
	for _, v := range nums2 {
		ans[1] += set1[v]
	}
	return ans[:]
}
