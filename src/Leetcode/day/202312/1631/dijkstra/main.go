package main

import (
	"container/heap"
	"math"
)

type Node struct{ x, y, v int }
type hp []Node

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].v < (*h)[j].v }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(Node)) }
func (h *hp) Pop() (v any)       { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func minimumEffortPath(heights [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n, m := len(heights), len(heights[0])
	dist := make([]int, m*n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}

	dist[0] = 0
	st := make([]bool, m*n)
	h := &hp{{0, 0, 0}}
	for h.Len() > 0 {
		p := heap.Pop(h).(Node)
		id := p.x*m + p.y
		if st[id] {
			continue
		}
		if p.x == n-1 && p.y == m-1 {
			break
		}
		st[id] = true
		for _, d := range dirs {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m &&
				max(p.v, abs(heights[x][y]-heights[p.x][p.y])) < dist[x*m+y] {
				dist[x*m+y] = max(p.v, abs(heights[x][y]-heights[p.x][p.y]))
				heap.Push(h, Node{x, y, dist[x*m+y]})
			}
		}
	}

	return dist[n*m-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {}
