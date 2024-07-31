package main

import "slices"

/*
	给你一个二维整数数组 point ，其中 points[i] = [xi, yi] 表示二维平面内的一个点。同时给你一个整数 w 。你需要用矩形 覆盖所有 点。

	每个矩形的左下角在某个点 (x1, 0) 处，且右上角在某个点 (x2, y2) 处，其中 x1 <= x2 且 y2 >= 0 ，同时对于每个矩形都 必须 满足 x2 - x1 <= w 。

	如果一个点在矩形内或者在边上，我们说这个点被矩形覆盖了。

	请你在确保每个点都 至少 被一个矩形覆盖的前提下，最少 需要多少个矩形。

	注意：一个点可以被多个矩形覆盖。
*/

/*
	由于矩形的高没有限制，所以我们只需考虑点的横坐标。

	矩形越宽，覆盖的点越多，所以 x2 应该恰好等于 x1 + w

	算法如下：
	1. 把横坐标按照从小到大的顺序排序。
	2. 为了方便计算，假设第一个举行左边还有一个矩形，初始化 x2 = -1，以为所有的横坐标都是非负数。
	3. 遍历横坐标 x = points[i][0]，如果 x > x2，我们需要一个新的 x1 = x 的矩形，答案加一，然后把 x2 更新为 x+w。
	4. 遍历结束，返回答案。
*/

func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {
	slices.SortFunc(points, func(p, q []int) int { return p[0] - q[0] })
	x2 := -1
	for _, p := range points {
		if p[0] > x2 {
			ans++
			x2 = p[0] + w
		}
	}
	return
}
