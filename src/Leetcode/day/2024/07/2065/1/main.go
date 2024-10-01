package main

/*
	给你一张 无向 图，图中有 n 个节点，节点编号从 0 到 n - 1 （都包括）。同时给你一个下标从 0 开始的整数数组 values ，其中 values[i]
	是第 i 个节点的 价值 。同时给你一个下标从 0 开始的二维整数数组 edges ，其中 edges[j] = [uj, vj, timej] 表示节点 uj 和 vj 之间
	有一条需要 timej 秒才能通过的无向边。最后，给你一个整数 maxTime 。
	合法路径 指的是图中任意一条从节点 0 开始，最终回到节点 0 ，且花费的总时间 不超过 maxTime 秒的一条路径。你可以访问一个节点任意次。一
	条合法路径的 价值 定义为路径中 不同节点 的价值 之和 （每个节点的价值 至多 算入价值总和中一次）。
	请你返回一条合法路径的 最大 价值。
	注意：每个节点 至多 有 四条 边与之相连。
*/

// 暴搜

func maximalPathQuality(values []int, edges [][]int, maxTime int) (ans int) {
	n := len(values)
	type edge struct{ to, time int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, t})
		g[y] = append(g[y], edge{x, t})
	}

	vis := make([]bool, n)
	vis[0] = true
	var dfs func(int, int, int)
	dfs = func(x, sumTime, sumValue int) {
		if x == 0 {
			ans = max(ans, sumValue)
			// 注意这里没有 return，还可以继续走
		}
		for _, e := range g[x] {
			y, t := e.to, e.time
			if sumTime+t > maxTime {
				continue
			}
			if vis[y] {
				dfs(y, sumTime+t, sumValue)
			} else {
				vis[y] = true
				// 每个节点的价值最多算入价值总和中一次
				dfs(y, sumTime+t, sumValue+values[y])
				vis[y] = false // 恢复现场
			}
		}
	}
	dfs(0, 0, values[0])
	return
}
