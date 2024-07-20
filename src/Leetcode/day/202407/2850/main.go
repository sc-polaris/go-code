package main

import "math"

/*
	给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目。网格图中总共恰好有 9 个石头，一个格子里可能会有 多个 石头。

	每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。

	请你返回每个格子恰好有一个石头的 最少移动次数 。

	枚举全排列：
	由于所有移走的石子个数等于所有移入的石子个数（即 0 的个数），我们可以把移走的石子的坐标记录到列表 from 中（可能有重复的坐标），
	移入的石子的坐标记录到列表 to 中。这两个列表的长度是一样的。

	枚举 from 的所有排列，与 to 匹配，即累加从 from[i] 到 to[i] 的曼哈顿距离。
	所有距离之和的最小值就是答案。
*/

type pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
	var from, to []pair
	for i, row := range grid {
		for j, cnt := range row {
			if cnt > 1 {
				for k := 1; k < cnt; k++ {
					from = append(from, pair{i, j})
				}
			} else if cnt == 0 {
				to = append(to, pair{i, j})
			}
		}
	}

	ans := math.MaxInt
	permute(from, 0, func() {
		total := 0
		for i, f := range from {
			total += abs(f.x-to[i].x) + abs(f.y-to[i].y)
		}
		ans = min(ans, total)
	})
	return ans
}

func permute(a []pair, i int, do func()) {
	if i == len(a) {
		do()
		return
	}
	permute(a, i+1, do)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permute(a, i+1, do)
		a[i], a[j] = a[j], a[i]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
