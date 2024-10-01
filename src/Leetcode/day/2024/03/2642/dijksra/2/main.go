package main

import (
	"container/heap"
	"math"
)

type Graph [][]pair

func Constructor(n int, edges [][]int) Graph {
	g := make(Graph, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], pair{e[1], e[2]})
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]] = append(g[e[0]], pair{e[1], e[2]})
}

func (g Graph) ShortestPath(start, end int) int {
	// dis[i] 表示从起点 start 出发，到节点 i 的最短路长度
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x, d := p.x, p.d
		if x == end { // 计算出从起点到终点的最短路长度
			return d
		}
		if d > dis[x] { // x 之前出堆过，无需更新邻居的最短路
			continue
		}
		for _, e := range g[x] {
			y, w := e.x, e.d
			newD := d + w
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{y, newD})
			}
		}
	}
	return -1 // 无法到达终点
}

type pair struct{ x, d int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; v = a[len(a)-1]; *h = a[:len(a)-1]; return }

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */

func main() {

}
