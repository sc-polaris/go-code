package dijkstra

import "container/heap"

/*
	建图：
	· x 到 ⌊x/2⌋ 连一条边权为 x mod 2 + 1 的 边。
	· x 到 ⌊x/3⌋ 连一条边权为 x mod 3 + 1 的 边。
	· 1 到 0 连一条边权为 1 的边。

	答案为 n 到 0 的最短路，用 dijkstra 算法计算。
*/

func minDays(n int) int {
	dis := make(map[int]int)
	h := &hp{{0, n}}
	for {
		p := heap.Pop(h).(pair)
		dx, x := p.d, p.x
		if x <= 1 {
			return dx + x
		}
		if dx > dis[x] {
			continue
		}
		for d := 2; d <= 3; d++ {
			y := x / d
			dy := dx + x%d + 1
			if dis[y] == 0 || dy < dis[y] {
				dis[y] = dy
				heap.Push(h, pair{dy, y})
			}
		}
	}
}

type pair struct{ d, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
