package main

import "math"

/*
	方法二：维护最大次大、最小次小
	优化：如果把最大的 x′ 移除，那么次大的 x′ 就是剩下 n−1 个 x′中的最大值了，对于最小值也同理。

	所以只需维护 x′ 和 y′ 的最大次大、最小次小，一共 8 个数。

	注意最大次大可能是相同的，最小次小可能是相同的。
*/

func minimumDistance(points [][]int) int {
	const inf = math.MaxInt
	maxX1, maxX2, maxY1, maxY2 := -inf, -inf, -inf, -inf
	minX1, minX2, minY1, minY2 := inf, inf, inf, inf

	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]

		// x 最大次大
		if x > maxX1 {
			maxX2, maxX1 = maxX1, x
		} else if x > maxX2 {
			maxX2 = x
		}

		// x 最小次小
		if x < minX1 {
			minX2, minX1 = minX1, x
		} else if x < minX2 {
			minX2 = x
		}

		// y 最大次大
		if y > maxY1 {
			maxY2, maxY1 = maxY1, y
		} else if y > maxY2 {
			maxY2 = y
		}

		// y 最小次小
		if y < minY1 {
			minY2, minY1 = minY1, y
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := inf
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		dx := f(x, maxX1, maxX2) - f(x, minX1, minX2)
		dy := f(y, maxY1, maxY2) - f(y, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(v, v1, v2 int) int {
	if v == v1 {
		return v2
	}
	return v1
}
