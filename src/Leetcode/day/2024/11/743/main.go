package main

import (
	"container/heap"
	"math"
	"slices"
)

/*
	有 n 个网络节点，标记为 1 到 n。

	给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。

	现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
*/

// dijkstra 稠密图
func networkDelayTime(times [][]int, n int, k int) int {
	const inf = math.MaxInt / 2
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		g[t[0]-1][t[1]-1] = t[2]
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[k-1] = 0
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done {
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 {
			return slices.Max(dis)
		}
		if dis[x] == inf { // 有节点无法到达
			return -1
		}
		done[x] = true // 最短路长度已确定（无法变得更小）
		for y, d := range g[x] {
			// 更新 x 的邻居的最短路
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
}

// 堆优化 Dijkstra（适用于稀疏图）
func networkDelayTime2(times [][]int, n int, k int) int {
	type edge struct{ to, wt int }
	g := make([][]edge, n) // 邻接表
	for _, t := range times {
		g[t[0]-1] = append(g[t[0]-1], edge{t[1] - 1, t[2]})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[k-1] = 0
	h := hp{{0, k - 1}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		dx, x := p.dis, p.x
		if dx > dis[x] { // x 之前出堆过
			continue
		}
		for _, e := range g[x] {
			y, newDis := e.to, dx+e.wt
			if newDis < dis[y] {
				dis[y] = newDis // 更新 x 的邻居的最短路
				heap.Push(&h, pair{newDis, y})
			}
		}
	}
	mx := slices.Max(dis)
	if mx < math.MaxInt {
		return mx
	}
	return -1
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
