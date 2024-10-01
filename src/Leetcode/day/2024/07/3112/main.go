package main

import "container/heap"

/*
	给你一个二维数组 edges 表示一个 n 个点的无向图，其中 edges[i] = [ui, vi, lengthi] 表示节点 ui 和节点 vi 之间有一条需要 lengthi 单位时间通过的无向边。
	同时给你一个数组 disappear ，其中 disappear[i] 表示节点 i 从图中消失的时间点，在那一刻及以后，你无法再访问这个节点。
	注意，图有可能一开始是不连通的，两个节点之间也可能有多条边。
	请你返回数组 answer ，answer[i] 表示从节点 0 到节点 i 需要的 最少 单位时间。如果从节点 0 出发 无法 到达节点 i ，那么 answer[i] 为 -1 。
*/

/*
	Dijkstra 算法

	对于本题，answer 几乎就是 dis 数组。只需要在 Dijkstra 算法的过程中，添加一个判断：
	· 在更新最短路之前，如果最短路长度 >= disappear[i]，说明无法及时到达节点 i，不更新。
*/

func minimumTime(n int, edges [][]int, disappear []int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n) // 稀疏图用邻接表
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	h := hp{{0, 0}}
	for h.Len() > 0 {
		p := heap.Pop(&h).(pair)
		dx, x := p.dis, p.x
		if dx > dis[x] { // 之前堆出现过
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newDis := dx + e.wt
			if newDis < disappear[y] && (dis[y] == -1 || newDis < dis[y]) {
				dis[y] = newDis // 更新 x 邻居的最短路
				heap.Push(&h, pair{newDis, y})
			}
		}
	}
	return dis
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
