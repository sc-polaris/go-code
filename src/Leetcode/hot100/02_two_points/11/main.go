package main

func maxArea(height []int) (ans int) {
	l, r := 0, len(height)-1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return
}
