package dijkstra

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	const INF = math.MaxInt32 >> 1
	g := make([][]int, n)  // 存储每条边
	dist := make([]int, n) // 存储0号点到每个点的最短距离
	st := make([]bool, n)  // 存储每个点的最短路是否已经确定

	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = INF
		}
	}
	for _, e := range edges {
		f, t, w := e[0], e[1], e[2]
		g[f][t], g[t][f] = w, w
	}

	dijkstra := func(u int) (cnt int) {
		for i := range st {
			st[i] = false
			dist[i] = INF
		}
		dist[u] = 0 // 自己到自己的距离为0
		for i := 0; i < n; i++ {
			k := -1 // 在还未确定最短路的点中，寻找距离最小的点
			for j := 0; j < n; j++ {
				if !st[j] && (k == -1 || dist[j] < dist[k]) {
					k = j
				}
			}

			// 用k更新其他点的距离
			for j := 0; j < n; j++ {
				dist[j] = min(dist[j], dist[k]+g[k][j])
			}
			st[k] = true
		}

		for _, v := range dist {
			if v <= distanceThreshold {
				cnt++
			}
		}
		return
	}

	ans, cnt := n, INF
	for i := n - 1; i >= 0; i-- {
		if t := dijkstra(i); t < cnt {
			cnt = t
			ans = i
		}
	}

	return ans
}
