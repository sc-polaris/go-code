package main

import "math"

/*
	一个公司在全国有 n 个分部，它们之间有的有道路连接。一开始，所有分部通过这些道路两两之间互相可以到达。
	公司意识到在分部之间旅行花费了太多时间，所以它们决定关闭一些分部（也可能不关闭任何分部），同时保证剩下的分部之间两两互相可以到达且最远距离不超过 maxDistance 。
	两个分部之间的 距离 是通过道路长度之和的 最小值 。
	给你整数 n ，maxDistance 和下标从 0 开始的二维整数数组 roads ，其中 roads[i] = [ui, vi, wi] 表示一条从 ui 到 vi 长度为 wi的 无向 道路。
	请你返回关闭分部的可行方案数目，满足每个方案里剩余分部之间的最远距离不超过 maxDistance。
	注意，关闭一个分部后，与之相连的所有道路不可通行。
	注意，两个分部之间可能会有多条道路。

	二进制枚举 + Floyd
	1. 根据 从集合论到位运算 中「枚举集合」的技巧，枚举 {0,1,2,⋯,n−1} 的所有子集 S，作为保留的节点集合，关闭（删除）不在 S 中的节点。
	2. 然后用 Floyd 算法求出任意两点之间的最短路。如果保留的节点之间的最短路均不超过 maxDistance，则把答案加一。

*/

func numberOfSets(n int, maxDistance int, roads [][]int) (ans int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt / 2 // 防止加法溢出
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	check := func(s int) int {
		for i, row := range g {
			if s>>i&1 == 1 {
				copy(f[i], row)
			}
		}

		// Floyd 算法（只考虑在 s 中的节点）
		for k := range f {
			if s>>k&1 == 1 {
				for i := range f {
					if s>>i&1 == 1 {
						for j := range f {
							f[i][j] = min(f[i][j], f[i][k]+f[k][j])
						}
					}
				}
			}
		}

		// 判断保留的节点之间的最短路是否均不超过 maxDistance
		for i, di := range f {
			if s>>i&1 == 1 {
				for j, dij := range di[:i] {
					if s>>j&1 == 1 && dij > maxDistance {
						return 0
					}
				}
			}
		}
		return 1
	}

	for s := 0; s < 1<<n; s++ { // 枚举子集
		ans += check(s)
	}
	return
}
