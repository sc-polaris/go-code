package main

func maxArea(height []int) (ans int) {
	l, r := 0, len(height)-1
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return
}
