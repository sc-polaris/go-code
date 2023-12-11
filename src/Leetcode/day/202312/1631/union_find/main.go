package main

import "sort"

type UnionFind struct {
	parent, size []int
}

func newUnionFind(n int) *UnionFind {
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
}

func (uf *UnionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

type Edge struct {
	x, y, v int
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	var edges []*Edge
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, &Edge{id - m, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, &Edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].v < edges[j].v })

	uf := newUnionFind(n * m)
	for _, e := range edges {
		uf.union(e.x, e.y)
		if uf.inSameSet(0, n*m-1) {
			return e.v
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

}
