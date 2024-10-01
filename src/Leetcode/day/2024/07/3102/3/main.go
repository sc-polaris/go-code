package main

import "math"

/*
	进一步地，要移除的点只能是 x′ 或 y′ 最大最小的点（不然移除前后都一样），所以额外维护最大最小值的下标，一共 4 个，
	最后只需遍历这 4 个坐标，而不是遍历整个 points 数组。
*/

func minimumDistance(points [][]int) int {
	const inf = math.MaxInt
	maxX1, maxX2, maxY1, maxY2 := -inf, -inf, -inf, -inf
	minX1, minX2, minY1, minY2 := inf, inf, inf, inf
	var maxXi, minXi, maxYi, minYi int

	for i, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]

		// x 最大次大
		if x > maxX1 {
			maxX2, maxX1, maxXi = maxX1, x, i
		} else if x > maxX2 {
			maxX2 = x
		}

		// x 最小次小
		if x < minX1 {
			minX2, minX1, minXi = minX1, x, i
		} else if x < minX2 {
			minX2 = x
		}

		// y 最大次大
		if y > maxY1 {
			maxY2, maxY1, maxYi = maxY1, y, i
		} else if y > maxY2 {
			maxY2 = y
		}

		// y 最小次小
		if y < minY1 {
			minY2, minY1, minYi = minY1, y, i
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := inf
	for _, i := range []int{maxXi, minXi, maxYi, minYi} {
		dx := f(i != maxXi, maxX1, maxX2) - f(i != minXi, minX1, minX2)
		dy := f(i != maxYi, maxY1, maxY2) - f(i != minYi, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(b bool, v1, v2 int) int {
	if b {
		return v1
	}
	return v2
}
