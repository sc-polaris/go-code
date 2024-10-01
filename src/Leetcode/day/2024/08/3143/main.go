package main

import (
	"math"
	"math/bits"
	"sort"
)

/*
	给你一个二维数组 points 和一个字符串 s ，其中 points[i] 表示第 i 个点的坐标，s[i] 表示第 i 个点的 标签 。

	如果一个正方形的中心在 (0, 0) ，所有边都平行于坐标轴，且正方形内 不 存在标签相同的两个点，那么我们称这个正方形是 合法 的。

	请你返回 合法 正方形中可以包含的 最多 点数。

	注意：
	· 如果一个点位于正方形的边上或者在边以内，则认为该点位于正方形内。
	· 正方形的边长可以为零。

	方法一：二分

	方法二：维护次小距离的最小值
	定义点 (x,y) 到 (0,0) 的切比雪夫距离为
							max(∣x∣,∣y∣)
	定义 minD[c]  为标签为 c 的所有点到 (0,0) 的最小切比雪夫距离。
	定义 minD2[c] 为标签为 c 的所有点到 (0,0) 的次小切比雪夫距离。

	那么正方形不能包含切比雪夫距离大于等于 min2 = min(minD2) 的点，否则正方形会包含标签相同的点。
	换句话说，可以包含的点需要满足
							minD[c] < min2
	代码实现时，无需维护 minD2[c]，而是直接维护 min2
*/

func maxPointsInsideSquare(points [][]int, s string) (ans int) {
	sort.Search(1_000_000_001, func(size int) bool {
		vis := 0
		for i, p := range points {
			if abs(p[0]) <= size && abs(p[1]) <= size { // 点在正方形内
				c := s[i] - 'a'
				if vis>>c&1 == 1 { // 在集合中
					return true
				}
				vis |= 1 << c // 把 c 加入集合
			}
		}
		ans = bits.OnesCount(uint(vis))
		return false
	})
	return
}

func maxPointsInsideSquare2(points [][]int, s string) (ans int) {
	minD := [26]int{}
	for i := range minD {
		minD[i] = math.MaxInt
	}
	min2 := math.MaxInt
	for i, p := range points {
		d := max(abs(p[0]), abs(p[1]))
		c := s[i] - 'a'
		if d < minD[c] {
			// d 目前是最小的，那么 minD[c] 是次小的
			min2 = min(min2, minD[c])
			minD[c] = d
		} else {
			// d 可能是次小的
			min2 = min(min2, d)
		}
	}

	for _, d := range minD {
		if d < min2 {
			ans++
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
