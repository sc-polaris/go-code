package main

import (
	"container/heap"
	"math"
)

/*
	堆优化版 dijkstra (适用于稀疏图)
	时间复杂度: O(mlogm) 空间复杂度: O(m) m 是 roads 的长度
*/

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func countPaths(n int, roads [][]int) int {
	type edge struct{ to, d int }
	g := make([][]edge, n) // 邻接表
	for _, r := range roads {
		x, y, d := r[0], r[1], r[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}
	dis := make([]int, n) // dis[i] 表示节点 0 到节点 i 的最短路长度
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	f := make([]int, n) // f[i] 表示节点 0 到节点 i 的最短路个数
	f[0] = 1
	h := &hp{{0, 0}}
	for {
		p := heap.Pop(h).(pair)
		x := p.x
		if x == n-1 {
			// 不可能找到比 dis[n-1] 更短，或者一样短的最短路了（注意本题边权都是正数）
			return f[n-1]
		}
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] { // 尝试更新 x 的邻居的最短路
			y := e.to
			newDis := p.dis + e.d
			if newDis < dis[y] {
				// 就目前来说，最短路一定经过 x
				dis[y] = newDis
				f[y] = f[x]
				heap.Push(h, pair{newDis, y})
			} else if newDis == dis[y] {
				f[y] = (f[y] + f[x]) % 1_000_000_007
			}
		}
	}
}

func main() {

}
