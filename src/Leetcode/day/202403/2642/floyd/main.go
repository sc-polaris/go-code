package main

import "math"

const inf = math.MaxInt / 3 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	f := make(Graph, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			if j != i {
				f[i][j] = inf
			}
		}
	}
	for _, e := range edges {
		f[e[0]][e[1]] = e[2]
	}
	for k := range f {
		for i := range f {
			if f[i][k] == inf {
				continue
			}
			for j := range f {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j])
			}
		}
	}
	return f
}

func (f Graph) AddEdge(e []int) {
	x, y, w := e[0], e[1], e[2]
	if w > f[x][y] { // 无需更新
		return
	}
	for i := range f {
		for j := range f {
			f[i][j] = min(f[i][j], f[i][x]+w+f[y][j])
		}
	}
}

func (f Graph) ShortestPath(start, end int) int {
	ans := f[start][end]
	if ans == inf {
		return -1
	}
	return ans
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */

func main() {

}
