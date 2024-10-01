package main

import "slices"

/*
	暴力做法是，枚举删除 x=initial[i]，然后从其它没有被删除的 initial[j] 开始，跑 DFS（或者 BFS 等），计算有多少个点
	能被感染到。对于每个 x，需要 O(n^2) 的时间遍历 graph，总的时间复杂度为 O(mn^2)，其中 m 为 initial 的长度。

	能否像 924 题那样做到 O(n^2) 时间？

	逆向思维，从不在 initial 中的点 v 出发 DFS，在不经过 initial 中的点的前提下，看看 v 是只能被一个点感染袋，还是能被
	多个点感染到。如果 v 只能被点 x == initial[i] 感染到，那么在本次 DFS 过程中访问到的其它节点，也只能被点 x 感染到。

	举例说明：
	initial = [1,5]
	0->[1,3]
	1->[0,2,4,5]
	2->[1,5]
	3->[0,4]
	4->[1,3]
	5->[1,2]
	我们从 0 出发 DFS，在不经过 1，5 的前提下，可以访问到节点 0，3，4，并且发现了恰好一个在 initial 中的节点 1。则意味着
	删除节点 1，就会让 0，3，4 免受感染。而从 2 出发 DFS，可以发现两个在 initial 中的节点，所以无论删除节点 1 还是删除节
	点 5，节点 2 一定会被感染。

	具体算法如下：
	1. 创建一个 vis 数组，标记在 DFS 中访问过的点。
	2. 枚举 [0,n-1] 中没有访问过的，且不在 initial 中的点 i。
	3. 从 i 开始 DFS。
	4. DFS 过程中，只访问不在 initial 中的节点，统计访问到的节点个数 size。
	5. DFS 过程中，如果发现了在 initial 中的节点，按照 924 中的状态集，更新变量 nodeId。
	6. DFS 结束后，如果 nodeId >= 0，那么把 nodeId（作为 key）和 size（作为 value）添加到一个哈希表或数组 cnt 中，其中
	   相同的 nodeId 要累加 size。
	7. 最后，如果 cnt 为空，返回 min(initial)；否则返回 cnt 中 size 最大的 nodeId，如果有多个 size 一样大，返回其中 nodeId
	   的最小值。

	代码实现时，可以用一个哈希表或者布尔数组，记录哪些点在 initial 中，从而在 DFS 中快速判断当前节点的邻居是否在 initial 中。
*/

func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	vis := make([]bool, n)
	isInitial := make([]bool, n)
	for _, x := range initial {
		isInitial[x] = true
	}

	var nodeId, size int
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		size++
		for y, conn := range graph[x] {
			if conn == 0 {
				continue
			}
			if isInitial[y] {
				// 按照 924 的状态机更新 nodeId
				if nodeId != -2 && nodeId != y {
					if nodeId == -1 {
						nodeId = y
					} else {
						nodeId = -2
					}
				}
			} else if !vis[y] {
				dfs(y)
			}
		}
	}
	cnt := make([]int, n)
	for i, seen := range vis {
		if seen || isInitial[i] {
			continue
		}
		nodeId = -1
		size = 0
		dfs(i)
		if nodeId >= 0 { // 只找到一个在 initial 中的节点
			// 删除节点 nodeId 可以让 size 个点不被感染
			cnt[nodeId] += size
		}
	}

	maxCnt := 0
	minNodeId := -1
	for i, c := range cnt {
		if c > maxCnt {
			maxCnt = c
			minNodeId = i
		}
	}
	if minNodeId >= 0 {
		return minNodeId
	}
	return slices.Min(initial)
}
