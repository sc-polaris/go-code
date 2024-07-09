package main

import (
	"github.com/emirpasic/gods/v2/trees/redblacktree"
	"math"
)

/*
	给你一个下标从 0 开始的数组 points ，它表示二维平面上一些点的整数坐标，其中 points[i] = [xi, yi] 。

	两点之间的距离定义为它们的 曼哈顿距离。

	请你恰好移除一个点，返回移除后任意两点之间的 最大 距离可能的 最小 值。
*/

/*
	https://leetcode.cn/problems/minimize-manhattan-distances/solutions/2716755/tu-jie-man-ha-dun-ju-chi-heng-deng-shi-b-op84/

	这两种投影长度，其中较大者为曼哈顿距离（较小者是两段折线的投影长度之差，不合法），即如下恒等式
								∣x1−x2∣+∣y1−y2∣=max(∣x1′−x2′∣,∣y1′−y2′∣)
	其中等式左侧为 (x1,y1) 和 (x2,y2) 的曼哈顿距离，等式右侧 (x′,y′)=(x+y,y−x)，计算的是 (x1′,y1′) 和 (x2′,y2′) 两点的曼哈顿距离投影到 x 轴
	和 y 轴的线段长度的最大值，即「切比雪夫距离」。

	所以要求任意两点曼哈顿距离的最大值，根据上面的恒等式，我们只需要计算任意两个 (x′,y′) 切比雪夫距离的最大值，即横纵坐标差的最大值
								max{max(x′)−min(x′),max(y′)−min(y′)}
*/

/*
	方法一 有序集合
	枚举要移除的点，用两个有序集合维护其他 n-1 个点的 x′ 和 y′，用 max{max(x′)−min(x′),max(y′)−min(y′)} 更新答案的最大值。

	// https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha
*/

func minimumDistance(points [][]int) int {
	xs := redblacktree.New[int, int]()
	ys := redblacktree.New[int, int]()
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		put(xs, x)
		put(ys, y)
	}

	ans := math.MaxInt
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		remove(xs, x) // 移除一个 x
		remove(ys, y) // 移除一个 y
		ans = min(ans, max(xs.Right().Key-xs.Left().Key, ys.Right().Key-ys.Left().Key))
		put(xs, x)
		put(ys, y)
	}
	return ans
}

func put(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	t.Put(v, c+1)
}

func remove(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	if c == 1 {
		t.Remove(v)
	} else {
		t.Put(v, c-1)
	}
}
