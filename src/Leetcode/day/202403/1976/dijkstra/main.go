package main

import "math"

/*
	朴素 dijkstra (适用于稠密图)
	时间复杂度: O(n^2) 空间复杂度: O(n^2)
*/

func countPaths(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt / 2 // 防止溢出
		}
	}
	for _, r := range roads {
		x, y, d := r[0], r[1], r[2]
		g[x][y] = d
		g[y][x] = d
	}

	dis := make([]int, n) // dis[i] 表示节点 0 到节点 i 的最短路长度
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt / 2
	}
	f := make([]int, n) // f[i] 表示节点 0 到节点 i 的最短路个数
	f[0] = 1
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x == n-1 {
			// 不可能找到比 dis[n-1] 更短，或者一样短的最短路了（本题边权都是正数）
			return f[n-1]
		}
		done[x] = true           // 最短路径已确定（无法变得更小）
		for y, d := range g[x] { // 尝试更新 x 的邻居的最短路
			newDis := dis[x] + d
			if newDis < dis[y] {
				// 就目前来说，最短路一定经过 x
				dis[y] = newDis
				f[y] = f[x]
			} else if newDis == dis[y] {
				f[y] = (f[y] + f[x]) % 1_000_000_007
			}
		}
	}
}

func main() {

}
