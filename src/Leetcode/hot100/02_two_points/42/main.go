package main

func trap(height []int) (ans int) {
	l, r, lMax, rMax := 0, len(height)-1, 0, 0
	for l < r {
		lMax = max(lMax, height[l])
		rMax = max(rMax, height[r])
		// 如果 height[l]<height[r]，则必有 lMax<rMax
		if lMax < rMax {
			ans += lMax - height[l]
			l++
		} else {
			ans += rMax - height[r]
			r--
		}
	}
	return
}
