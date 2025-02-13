package main

func longestConsecutive(nums []int) (ans int) {
	has := make(map[int]bool)
	for _, num := range nums {
		has[num] = true
	}
	for x := range has {
		if has[x-1] {
			continue
		}
		// x 是序列起点
		y := x + 1
		for has[y] {
			y++
		}
		ans = max(ans, y-x)
	}
	return
}
