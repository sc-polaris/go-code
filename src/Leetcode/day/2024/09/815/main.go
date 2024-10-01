package main

/*
	给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。

	例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... 这样的车站路线行驶。
	现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。

	求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。
*/

/*
	有哪些公交车会经过车站 x？
	创建一个哈希表 stopToBuses，key 为车站编号，value 为经过该车站的公交车编号列表。
	遍历第 i 辆公交车的路线 routes[i]，对于车站 x=routes[i][j]，把公交车编号 i 加到 stopToBuses[x] 列表中。

	在 BFS 中，如何保证每辆公交车的路线只遍历一次？
	可以创建一个 vis 数组。更简单的办法是，当公交车路线 routes[i] 遍历结束后，把 routes[i] 置为空。

	在 BFS 中，如何保证每个车站只入队一次？
	为了记录起点到每个站的最短路（最少乘坐的公交车数量），创建一个哈希表 dis，key 为车站编号，value 为起点到该车站的最短路。

	我们可以利用 dis 来知道车站 x 是否入队过：看 x 是否在 dis 中即可。

	小优化
	如果没有公交车经过起点或终点，直接返回：
	· 如果 source!=target，无法从起点到达终点，返回 −1。
	· 如果 source=target，返回 0。

*/

func numBusesToDestination(routes [][]int, source int, target int) int {
	// 记录经过车站 x 的公交车编号
	stopToBuses := make(map[int][]int)
	for i, route := range routes {
		for _, x := range route {
			stopToBuses[x] = append(stopToBuses[x], i)
		}
	}

	// 小油画：如果没有公交车经过起点或终点，直接返回
	if stopToBuses[source] == nil || stopToBuses[target] == nil {
		if source != target {
			return -1
		}
		// 原地 TP 的情况
		return 0
	}

	// BFS
	dis := map[int]int{source: 0}
	q := []int{source}
	for len(q) > 0 {
		x := q[0] // 当前在车站 x
		q = q[1:]
		disX := dis[x]
		for _, i := range stopToBuses[x] { // 遍历所有经过车站 x 的公交车 i
			for _, y := range routes[i] { // 遍历公交车 i 的路线
				if _, ok := dis[y]; !ok { // 没有访问过车站 y
					dis[y] = disX + 1 // 从 x 站上车然后在 y 站下车
					q = append(q, y)
				}
			}
			routes[i] = nil // 标记 routes[i] 遍历过
		}
	}

	if d, ok := dis[target]; ok {
		return d
	}
	return -1
}
