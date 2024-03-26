package main

import "math"

const inf = math.MaxInt / 2 // 防止更新最短路的时候加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make(Graph, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // 添加一条边 题目保证没有重边
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2] // 添加一条边（题目保证这条边之前不存在）
}

func (g Graph) ShortestPath(start, end int) int {
	n := len(g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在为无穷大
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == end { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true //  标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路径
		}
	}
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */

func main() {

}
