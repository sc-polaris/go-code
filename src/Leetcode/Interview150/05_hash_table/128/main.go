package main

import "slices"

func longestConsecutive(nums []int) (ans int) {
	slices.Sort(nums)
	m := make(map[int]bool)
	length := 1
	for _, x := range nums {
		if m[x] {
			continue
		}
		m[x] = true
		if !m[x-1] {
			length = 1
		} else {
			length++
		}
		ans = max(ans, length)
	}
	return
}
